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
	"strings"
)

type server struct {
	apikeyserver.UnimplementedApiKeyServerServer
}

// TODO: pass pointer to GetKeyResponse to next()
func (s *server) GetKey(ctx context.Context, requester *apikeyserver.RequestKey) (*apikeyserver.GetKeyResponse, error) {
	Log.Debug().Caller().Msg("GetKey()")
	reqStr := requester.Requester
	reqType := requester.Type
	acceptExhaustion := requester.AcceptExhaustion
	Log.Info().Str("requester", reqStr).Str("type", reqType).Msg("Received request")
	responseStruct := next(&keys, reqType, acceptExhaustion)

	// TODO: extract to method that can work with KeyResponseRemaining and KeyDetailsRespons
	var keysLeft []*apikeyserver.KeyResponseRemaining

	mutexKeys.Lock()
	for _, v := range keys.Apikeys {
		types := strings.Join(v.Types, ", ")
		key := &apikeyserver.KeyResponseRemaining{
			KeyResponseTypeNames: types,
			TypeRemaining:        uint32(v.CurrentlyRemaining),
		}
		keysLeft = append(keysLeft, key)
	}
	mutexKeys.Unlock()
	// end extract method

	return &apikeyserver.GetKeyResponse{
		Key:       responseStruct.key,
		Name:      responseStruct.name,
		Type:      responseStruct.keyType,
		Time:      responseStruct.time,
		Exhausted: responseStruct.exhausted,
		Items:     keysLeft,
	}, nil
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

func (s *server) GetServerInfo(ctx context.Context, request *apikeyserver.RequestServerInfo) (*apikeyserver.GetServerInfoResponse, error) {
	Log.Debug().Caller().Msg("GetServerInfo()")
	Log.Info().Str("requester", request.Requester).Str("fieldmask", request.FieldMask.String()).Msg("Server Info")

	res := apikeyserver.GetServerInfoResponse{}

	return collectServerInfo(&keys, request, &res), nil
}
