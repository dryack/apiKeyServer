// apiKeyServer - gRCP/protobufs API key server for Torn API
//    Copyright (C) 2021  Dave Ryack
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU Affero General Public License as published
//    by the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU Affero General Public License for more details.
//
//    You should have received a copy of the GNU Affero General Public License
//    along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	pb "apiKeyServer/apikeyserver"
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tebeka/atexit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"text/tabwriter"
	"time"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	// port       	= flag.Int("port", 50051, "The server port")
	port          = flag.Int("port", 50052, "The server port") // for remote debugging purposes
	keys          Keys
	exhausted     = 0
	t             int64
	serverVersion = "v1.24"
	// Log setting up the logger object for global access
	Log     zerolog.Logger
	Sampled zerolog.Logger
	// Set up ticker and teardown of go routine
	Done   = make(chan bool)                        // necessary because we can't pass args to exitHandler()
	Ticker = time.NewTicker(250 * time.Millisecond) // set up ticker for checking the minute
	// lock keys when needed
	mutexKeys = sync.RWMutex{}
)

// TODO: Options for logging; turn it off, change logfile location, etc.
func init() {
	consoleLogging := flag.Bool("console", false, "display to console in addition to log")
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	logfile, err := os.OpenFile("./apikeyserver.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		Sampled = log.Sample(&zerolog.BasicSampler{N: 10})
	}

	if *consoleLogging {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		multi := zerolog.MultiLevelWriter(consoleWriter, logfile)
		Log = zerolog.New(multi).With().Timestamp().Logger()
	} else {
		Log = zerolog.New(logfile).With().Timestamp().Logger()
	}
	Log.Debug().Caller().Msg("completed: init()")
}

func main() {
	tWriter := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', 0)
	file, err := ioutil.ReadFile("./configs/config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &keys)
	if err != nil {
		panic(err)
	}

	t = time.Now().UTC().UnixMilli() + 60000 // 1 minute
	initKeys(&keys)

	Log.Info().Msg("Torn API Key server " + serverVersion)
	fmt.Println("lamashtu's Torn API Key server " + serverVersion)
	Log.Info().
		Str("max keys/min", strconv.Itoa(keys.TotalPerMinute)).
		Str("keys available", strconv.Itoa(len(keys.Apikeys))).
		Msg("")
	_, _ = fmt.Fprintf(os.Stdout, "%v keys available for use, up to %v queries per minute\n", len(keys.Apikeys), keys.TotalPerMinute)
	for k := range keys.Apikeys {
		_, _ = fmt.Fprintf(tWriter, "%s\t%v\t%s\t%s%s\n", keys.Apikeys[k].User, keys.Apikeys[k].MaxPerMinute, " uses/min", " types: ", keys.Apikeys[k].Types)
		Log.Info().
			Str("keyUser", keys.Apikeys[k].User).
			Str("keyMaxUsers", strconv.Itoa(keys.Apikeys[k].MaxPerMinute)).
			Str("types", strings.Join(keys.Apikeys[k].Types, ",")).
			Msg("")
	}
	err = tWriter.Flush() // sends column-formatted output to stdio
	if err != nil {
		panic(err)
	}

	atexit.Register(exitHandler)

	//flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		Log.Error().
			Err(err).
			Msg("failed to listen")

	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			Log.Error().
				Err(err).
				Msg("Failed to generate credentials")
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	// intercept SIGINT --
	// https://stackoverflow.com/questions/11268943/is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in-a-defe
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		exitHandler()
		os.Exit(0)
	}()

	// set up server
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterApiKeyServerServer(grpcServer, &server{})
	// start serving
	fmt.Printf("Now serving keys on port %v\n", *port)
	Log.Info().
		Str("port", strconv.Itoa(*port)).
		Msg("Now serving keys")

	// watching the time so we can re-init keys each minute
	go checkMinute(&keys)

	_ = grpcServer.Serve(lis)

	atexit.Exit(0)
}
