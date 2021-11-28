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
	"github.com/mennanov/fmutils"
)

type server struct {
	apikeyserver.UnimplementedApiKeyServerServer
}

// TODO: pass pointer to GetKeyResponse to next()
func (s *server) GetKey(ctx context.Context, request *apikeyserver.RequestKey) (*apikeyserver.GetKeyResponse, error) {
	Log.Debug().Caller().Msg("GetKey()")
	reqStr := request.Requester
	reqType := request.Type
	acceptExhaustion := request.AcceptExhaustion
	Log.Info().Str("request", reqStr).Str("type", reqType).Msg("Received request")
	res := next(&keys, reqType, acceptExhaustion)

	fmutils.Filter(res, request.FieldMask.GetPaths())
	return res, nil
}

func (s *server) KillKey(ctx context.Context, request *apikeyserver.RequestKillKey) (*apikeyserver.GenericKillResponse, error) {
	Log.Debug().Caller().Msg("KillKey()")
	keyToKill := request.Key
	Log.Info().Str("request", keyToKill).Msg("Killing ")
	killKey(&keys, keyToKill)
	return &apikeyserver.GenericKillResponse{Result: true}, nil
}

func (s *server) PermKillKey(ctx context.Context, request *apikeyserver.RequestPermKillKey) (*apikeyserver.GenericKillResponse, error) {
	Log.Debug().Caller().Msg("PermKillKey()")
	keyToKill := request.Key
	Log.Info().Str("request", keyToKill).Msg("Permanently killing ")
	permKillKey(&keys, keyToKill)
	return &apikeyserver.GenericKillResponse{Result: true}, nil
}

func (s *server) GetServerInfo(ctx context.Context, request *apikeyserver.RequestServerInfo) (*apikeyserver.GetServerInfoResponse, error) {
	Log.Debug().Caller().Msg("GetServerInfo()")
	Log.Info().Str("requester", request.Requester).Str("fieldmask", request.FieldMask.String()).Msg("Server Info")

	return collectServerInfo(&keys, request), nil
}
