package app

import (
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/stake"

	abci "github.com/tendermint/tendermint/abci/types"
)

func setGenesis(app *ThorchainApp, accs ...*auth.BaseAccount) error {
	genaccs := make([]GenesisAccount, len(accs))
	for i, acc := range accs {
		genaccs[i] = NewGenesisAccount(acc)
	}

	genesisState := GenesisState{
		Accounts:  genaccs,
		StakeData: stake.DefaultGenesisState(),
	}

	stateBytes, err := wire.MarshalJSONIndent(app.cdc, genesisState)
	if err != nil {
		return err
	}

	// Initialize the chain
	vals := []abci.Validator{}
	app.InitChain(abci.RequestInitChain{Validators: vals, AppStateBytes: stateBytes})
	app.Commit()

	return nil
}
