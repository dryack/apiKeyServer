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

// functions related to time

package main

import "time"

func getTime() int64 {
	Log.Debug().Caller().Msg("getTime()")
	newTime := time.Now().UTC().Unix()
	return newTime
}

// TODO this can probably be better (and more accurately) accomplished using https://gobyexample.com/tickers
// if a minute has passed, we can reset CurrentlyRemaining in Keys
func checkMinute(keys *Keys) {
	Log.Debug().Caller().Msg("checkMinute()")
	newT := getTime()
	if newT > t {
		t = newT + 60
		initKeys(keys)
	}
}
