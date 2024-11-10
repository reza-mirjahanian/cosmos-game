package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	checkersv1 "github.com/alice/checkers/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: checkersv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GetGame",
					Use:       "get-game index",
					Short:     "Get the current value of the game at index",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "index"},
					},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: checkersv1.CheckersTorram_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CheckersCreateGm",
					Use:       "create index black red",
					Short:     "Creates a new checkers game at the index for the black and red players",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "index"},
						{ProtoField: "black"},
						{ProtoField: "red"},
					},
				},
				{
					RpcMethod: "KillGame",
					Use:       "kill-game stop and store finish time",
					Short:     "Kills the game at the index",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "index"},
					},
				},
			},
		},
	}
}
