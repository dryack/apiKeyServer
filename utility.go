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

// utility functions needed elsewhere

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

// stolen from: https://gist.github.com/NorbertFenk/7bed6760198800207e84f141c41d93c7
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func exitHandler() {
	Log.Info().Msg("Exiting by user command.")
	Log.Debug().Caller().Msg("exitHandler()")
	Done <- true
	fmt.Print("\n\n")
	for _, key := range keys.Apikeys {
		Log.Info().
			Str("key", key.Tornkey).
			Str("kills", strconv.FormatUint(uint64(key.Kills), 10)).
			Str("uses", strconv.FormatUint(uint64(key.Uses), 10)).
			Msg("")
	}
	Log.Info().Str("exhaustions", strconv.Itoa(int(keys.TotalExhaustions))).Msg("")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	Log.Debug().Msgf("%s took %s", name, elapsed)
}

func startMessages(tabWriter *tabwriter.Writer, err error) error {
	Log.Info().Msg("Torn API Key server " + serverVersion)
	fmt.Println("lamashtu's Torn API Key server " + serverVersion)
	Log.Info().
		Str("max keys/min", strconv.Itoa(int(keys.TotalPerMinute))).
		Str("keys available", strconv.Itoa(len(keys.Apikeys))).
		Msg("")
	_, _ = fmt.Fprintf(os.Stdout, "%v keys available for use, up to %v queries per minute\n", len(keys.Apikeys), keys.TotalPerMinute)
	for k := range keys.Apikeys {
		_, _ = fmt.Fprintf(tabWriter, "%s\t%v\t%s\t%s%s\n", keys.Apikeys[k].User, keys.Apikeys[k].MaxPerMinute, " uses/min", " types: ", keys.Apikeys[k].Types)
		Log.Info().
			Str("keyUser", keys.Apikeys[k].User).
			Str("keyMaxUsers", strconv.Itoa(int(keys.Apikeys[k].MaxPerMinute))).
			Str("types", strings.Join(keys.Apikeys[k].Types, ",")).
			Msg("")
	}
	err = tabWriter.Flush() // sends column-formatted output to stdio
	if err != nil {
		panic(err)
	}
	return err
}
