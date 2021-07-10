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
	"strconv"
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
	fmt.Print("\n\n")
	for _, key := range keys.Apikeys {
		Log.Info().
			Str("key", key.Tornkey).
			Str("kills", strconv.FormatUint(key.Kills, 10)).
			Str("uses", strconv.FormatUint(key.Uses, 10)).
			Msg("")
		//fmt.Print(key.Tornkey + ": " + strconv.FormatUint(key.Kills, 10) + "\n")
	}
	Log.Info().Str("exhaustions", strconv.Itoa(exhausted)).Msg("")
	//fmt.Print("Number of exhaustions: " + strconv.Itoa(exhausted) + "\n")
}
