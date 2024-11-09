package keeper

import (
	"context"
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/collections"
	"github.com/alice/checkers"
	"github.com/alice/checkers/rules"
)

type checkersTorramServer struct {
	k Keeper
}

var _ checkers.CheckersTorramServer = checkersTorramServer{}

// NewMsgServerImpl returns a message server implementation
func NewMsgServerImpl(keeper Keeper) checkers.CheckersTorramServer {
	return &checkersTorramServer{
		k: keeper,
	}
}

// CheckersCreateGm creates a new checkers game
func (ms checkersTorramServer) CheckersCreateGm(ctx context.Context, msg *checkers.ReqCheckersTorram) (*checkers.ResCheckersTorram, error) {
	fmt.Printf("KillGame (msg) is: %v\n", msg)
	if length := len([]byte(msg.Index)); checkers.MaxIndexLength < length || length < 1 {
		return nil, checkers.ErrIndexTooLong
	}
	if _, err := ms.k.StoredGames.Get(ctx, msg.Index); err == nil || errors.Is(err, collections.ErrEncoding) {
		return nil, fmt.Errorf("game already exists at index: %s", msg.Index)
	}

	unwrappedCtx := sdk.UnwrapSDKContext(ctx)

	newBoard := rules.New()
	storedGame := checkers.StoredGame{
		Board:     newBoard.String(),
		Turn:      rules.PieceStrings[newBoard.Turn],
		Black:     msg.Black,
		Red:       msg.Red,
		StartTime: unwrappedCtx.BlockTime().UTC().String(),
		EndTime:   "",
	}

	if err := storedGame.Validate(); err != nil {
		return nil, err
	}
	if err := ms.k.StoredGames.Set(ctx, msg.Index, storedGame); err != nil {
		return nil, err
	}

	return &checkers.ResCheckersTorram{}, nil
}

// todo logic is not implemented completely
// a - who allow to kill the game
// b - what is the effect of killing the game on the game state
// c - when is possible to kill the game
func (ms checkersTorramServer) KillGame(ctx context.Context, msg *checkers.ReqKillGame) (*checkers.ResKillGame, error) {
	unwrappedCtx := sdk.UnwrapSDKContext(ctx)
	game, err := ms.k.StoredGames.Get(unwrappedCtx, msg.Index)
	if err != nil {
		fmt.Printf("Game not found! Error: %v\n", err)
		return nil, err
	}

	if game.EndTime != "" {
		fmt.Printf("Game already ended! \n")
		return nil, errors.New("game already ended")
	}

	//  Save the end time in state
	game.EndTime = unwrappedCtx.BlockTime().UTC().String()

	//  Save the game state in state
	if err := ms.k.StoredGames.Set(unwrappedCtx, msg.Index, game); err != nil {
		return nil, err
	}

	return &checkers.ResKillGame{}, nil
}
