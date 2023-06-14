package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

	"seshat/x/seshat/types"
)

//just keeper module can call this function
//to use the types here, need to transpile the proto to go
//ignite generate proto-go
func (k Keeper) SayHello(goCtx context.Context, req *types.QuerySayHelloRequest) (*types.QuerySayHelloResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	//process the query
	_ = ctx
	return &types.QuerySayHelloResponse{Name: fmt.Sprintf("hello %s", req.Name)}, nil
}