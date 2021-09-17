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

// boilerplate for apiKeyServer gRPC

package main

import (
	"apiKeyServer/apikeyserver"
	"context"
)

type server struct {
	apikeyserver.UnimplementedApiKeyServerServer
}

func (s *server) GetKey(ctx context.Context, requester *apikeyserver.RequestKey) (*apikeyserver.GetKeyResponse, error) {
	Log.Debug().Caller().Msg("GetKey()")
	reqStr := requester.Requester
	reqType := requester.Type
	Log.Info().Str("requester", reqStr).Str("type", reqType).Msg("Received request")
	key, name := next(&keys, reqType)
	return &apikeyserver.GetKeyResponse{Key: key, Name: name}, nil
}

func (s *server) KillKey(ctx context.Context, key *apikeyserver.RequestKillKey) (*apikeyserver.GenericKillResponse, error) {
	Log.Debug().Caller().Msg("KillKey()")
	keyToKill := key.Key
	Log.Info().Str("key", keyToKill).Msg("Killing ")
	killKey(&keys, keyToKill)
	return &apikeyserver.GenericKillResponse{Result: true}, nil
}

func (s *server) PermKillKey(ctx context.Context, key *apikeyserver.PermRequestKillKey) (*apikeyserver.GenericKillResponse, error) {
	Log.Debug().Caller().Msg("PermKillKey()")
	keyToKill := key.Key
	Log.Info().Str("key", keyToKill).Msg("Permanently killing ")
	permKillKey(&keys, keyToKill)
	return &apikeyserver.GenericKillResponse{Result: true}, nil
}
