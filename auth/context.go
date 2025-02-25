package auth

import (
	"context"

	"google.golang.org/grpc/metadata"

	"github.com/piotrostr/atlas-app-toolkit/v2/gateway"
	"github.com/piotrostr/atlas-app-toolkit/v2/requestid"
)

// OutgoingContext set to outgoing context request_id, auth_token, X-Forwarded-For, x-geo- and x-b3- headers value
func OutgoingContext(ctx context.Context) context.Context {
	keys := []string{
		AuthorizationHeader,
		requestid.DefaultRequestIDKey,
		gateway.XForwardedFor,
	}

	keys = append(keys, gateway.GetGeoHeaders()...)
	keys = append(keys, gateway.GetXB3Headers()...)

	return OutgoingContextWithCustomParams(ctx, keys...)
}

// OutgoingContextWithCustomParams set to outgoing context specified parameters from incoming context by keys
func OutgoingContextWithCustomParams(ctx context.Context, keys ...string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	resultMD := make(metadata.MD, 0)

	for _, key := range keys {
		if params := md.Get(key); params != nil {
			resultMD.Append(key, params...)
		}
	}

	return metadata.NewOutgoingContext(ctx, resultMD)
}
