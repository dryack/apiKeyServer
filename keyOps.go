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

// definition of and operations against Keys struct

package main

import (
	"apiKeyServer/apikeyserver"
	"fmt"
	"github.com/mennanov/fmutils"
	"strconv"
	"strings"
	"time"
)

// TODO: (Optionally) return a message signalling keys are exhausted, this would permit the client program to continue
// processing. Currently, we just wait until keys again become available before returning a result. This forces the client
// to wait for a result or timeout. With an appropriate timeout, this has worked fine for me so far, but seems limiting as
// a more general purpose solution.

// TODO: Request multiple keys in a single request. It may be more efficient for the client to request X keys of type Y at
// once.

type Keys struct {
	TotalExhaustions uint32
	TotalPerMinute   uint32
	TotalKeysServed  uint64
	ServerVersion    string
	StartupTime      time.Time
	Apikeys          []struct {
		User               string   `yaml:"user"`
		MaxPerMinute       uint32   `yaml:"max_per_minute"`
		Tornkey            string   `yaml:"tornkey"`
		CurrentlyRemaining int32    `yaml:"currently_remaining"`
		Types              []string `yaml:"types"`
		Active             bool     `yaml:"active"`
		Kills              uint32
		Uses               uint32
	} `yaml:"apikeys"`
}

func initKeys(keys *Keys) {
	Log.Info().Msg("Initializing keys")
	Log.Debug().Caller().Msg("initKeys()")

	mutexKeys.Lock()
	defer mutexKeys.Unlock()

	if keys.StartupTime.UnixNano() == 0 {
		keys.StartupTime = time.Now()
	}

	if keys.ServerVersion == "" {
		keys.ServerVersion = serverVersion
	}

	keys.TotalPerMinute = 0
	for i := range keys.Apikeys {
		if keys.Apikeys[i].Active {
			keys.TotalPerMinute += keys.Apikeys[i].MaxPerMinute
		}
		keys.Apikeys[i].CurrentlyRemaining = int32(keys.Apikeys[i].MaxPerMinute)
	}
}

// any keys have uses remaining right now?
func anyLeft(keys Keys, keyType string) bool {
	defer timeTrack(time.Now(), "anyLeft()") // debug
	Log.Debug().Caller().Str("keyType", keyType).Msg("anyLeft()")
	for _, key := range keys.Apikeys {
		if contains(key.Types, keyType) {
			if key.Active {
				if key.CurrentlyRemaining > 0 {
					return true
				}
			}
		}
	}
	Log.Debug().Str("keyType", keyType).Msg("No keys found")
	return false
}

// TODO: can i use a heap-based priority queue and replace most of this logic - possibly even most of keyOps.go
func levelKeyUses(keys *Keys, keyType string) (string, string) {
	defer timeTrack(time.Now(), "levelKeyUses()") // debug
	Log.Debug().Caller().Str("keyType", keyType).Msg("levelKeyUses()")

	var max int32 = -1
	var contained = false
	for _, key := range keys.Apikeys {
		if contains(key.Types, keyType) && key.Active {
			contained = true
			if key.CurrentlyRemaining > 0 && key.CurrentlyRemaining > max {
				max = key.CurrentlyRemaining
			}
		}
	}
	for i, key := range keys.Apikeys {
		if contained && key.Active {
			if key.CurrentlyRemaining == max {
				keys.Apikeys[i].CurrentlyRemaining -= 1
				Log.Debug().
					Str("key", key.Tornkey).
					Msg("returning key ")
				keys.Apikeys[i].Uses++
				keys.TotalKeysServed++
				return key.Tornkey, key.User
			}
		}
	}
	Log.Debug().Msg("returning \"nil\"")
	return "nil", "nil"
}

// TODO: eliminate this struct, see comment at next() below
type getKeyResponse struct {
	key       string
	name      string
	keyType   string
	time      int64
	exhausted bool
}

// TODO: refactor to accept pointer to GetKeyResponse and replace type struct getKeyResponse usage
// return a key for use by requester
func next(keys *Keys, keyType string, acceptExhaustion bool) *getKeyResponse {
	Log.Debug().Caller().Str("keyType", keyType).Msg("next()")
	firstrun := 0

	for {
		Log.Debug().Str("runs", strconv.Itoa(firstrun)).Msg("next loop running")
		if anyLeft(*keys, keyType) {
			key, name := levelKeyUses(keys, keyType)
			// stolen from: https://stackoverflow.com/questions/16331063/how-can-i-get-the-string-representation-of-a-struct
			keysForSample := fmt.Sprintf("%v", keys.Apikeys)
			Sampled.Debug().Msg(keysForSample)

			return &getKeyResponse{
				name:      name,
				key:       key,
				keyType:   keyType, // returning the *requested* keytype - not the available keytypes for a given record
				time:      time.Now().UnixNano(),
				exhausted: false,
			}
		} else if acceptExhaustion {
			return &getKeyResponse{
				time:      time.Now().UnixNano(),
				exhausted: true,
			}
		}
		// we're out of keys.  if this is our first time at exhaustion this minute, print a message.  subsequent loops
		// where keys are still exhausted will be silent during non-debug operation
		if firstrun == 0 {
			Log.Info().Msg("Waiting for key to become available")
			keys.TotalExhaustions++
			Log.Debug().Caller().Str("exhaustion-cycle", strconv.Itoa(int(keys.TotalExhaustions))).Msg("Exhaustion")
		}
		firstrun++
		time.Sleep(1 * time.Second)
	}
}

func killKey(keys *Keys, keyToKill string) {
	defer timeTrack(time.Now(), "killKey()")
	Log.Debug().Caller().
		Str("key", keyToKill).
		Msg("killKey()")
	for i := range keys.Apikeys {
		if keys.Apikeys[i].Tornkey == keyToKill {
			keys.Apikeys[i].CurrentlyRemaining = 0
			keys.Apikeys[i].Kills += 1
		}
	}
}

func permKillKey(keys *Keys, keyToKill string) {
	defer timeTrack(time.Now(), "permKillKey()")
	Log.Debug().Caller().
		Str("key", keyToKill).
		Msg("permKillKey()")
	for i := range keys.Apikeys {
		if keys.Apikeys[i].Tornkey == keyToKill {
			keys.Apikeys[i].CurrentlyRemaining = 0
			keys.Apikeys[i].Kills += 1
			keys.Apikeys[i].Active = false
		}
	}
}

func collectServerInfo(keys *Keys, req *apikeyserver.RequestServerInfo, res *apikeyserver.GetServerInfoResponse) *apikeyserver.GetServerInfoResponse {
	uptime := time.Since(keys.StartupTime)

	var totKilled uint32
	var permKilled []string
	var keyDetails []*apikeyserver.KeyDetailsResponse
	mutexKeys.Lock()
	for _, v := range keys.Apikeys {
		types := strings.Join(v.Types, ", ")
		totKilled += v.Kills
		if !v.Active {
			permKilled = append(permKilled, v.User)
		}
		key := &apikeyserver.KeyDetailsResponse{
			Name:   v.User,
			Types:  types,
			Uses:   v.Uses,
			Kills:  v.Kills,
			Active: v.Active,
		}
		keyDetails = append(keyDetails, key)
	}
	mutexKeys.Unlock()

	res.ServerVersion = serverVersion
	res.KeyExhaustions = keys.TotalExhaustions
	res.TotalKeysServed = keys.TotalKeysServed
	res.TotalAvailableUsesPerMin = uint64(keys.TotalPerMinute)
	res.KeyNamesPermaKilled = strings.Join(permKilled, ", ")
	res.Items = keyDetails
	res.Time = time.Now().UnixNano()
	res.Uptime = int64(uptime)
	res.AvgKeysServedPerMin = float32((float64(keys.TotalKeysServed)) / uptime.Minutes())

	fmutils.Filter(res, req.FieldMask.GetPaths())
	return res
}
