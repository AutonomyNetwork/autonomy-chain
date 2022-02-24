package launchpad

import (
	"fmt"

	"github.com/AutonomyNetwork/autonomy-chain/x/launchpad/keeper"
	"github.com/AutonomyNetwork/autonomy-chain/x/launchpad/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func LaunchpadEndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	createdLaunchpads := k.GetCreatedLaunchpads(ctx)
	clps := createdLaunchpads.CreatedLaunchpads
	var clpsids []uint64
	var alpsids []uint64
	fmt.Println("Created launchpads............", clps)
	for _, l := range clps {
		fmt.Println("launchpad-id", l)
		lp := k.GetLaunchpad(ctx, l)
		fmt.Println("launchpad", lp)
		if lp.StartTime.Before(ctx.BlockTime()) {
			fmt.Println("if condition true", lp.StartTime, ctx.BlockTime())
			lp.Status = "ACTIVE"
			k.SetLaunchpad(ctx, lp)
			alpsids = append(alpsids, lp.Id) // Append to active lanchpads
		} else {
			clpsids = append(clpsids, lp.Id)
		}
	}

	fmt.Println("After set created launcpads", clpsids)
	fmt.Println("After set active launcpads", alpsids)

	activeLaunchpads := k.GetActiveLaunchpads(ctx)
	alps := activeLaunchpads.ActiveLaunchpads
	var alpsidss []uint64

	slps := k.GetSuccessLaunchpads(ctx).SuccessLaunchpads
	flps := k.GetFailLaunchpads(ctx).FailLaunchpads

	status := ""
	var slpids []uint64
	var flpids []uint64
	fmt.Println("Success launchpad ids", slps)
	fmt.Println("Failed launcpad ids", flps)
	fmt.Println("Current active launcpads", alps)
	for _, l := range alps {
		lp := k.GetLaunchpad(ctx, l)
		fmt.Println("Active launchpad", lp)
		if lp.EndTime.Before(ctx.BlockTime()) {
			fmt.Println("If condition true")
			if lp.Deposits >= lp.Softcap && lp.Deposits <= lp.Hardcap {
				status = "SUCCESS"
				slpids = append(slpids, lp.Id)
				//TODO : create vesting account
			} else {
				status = "FAIL"
				flpids = append(flpids, lp.Id)
			}

			lp.Status = status
			k.SetLaunchpad(ctx, lp)
		} else {
			alpsidss = append(alpsidss, lp.Id)
		}
	}

	fmt.Println("After set Total active launchpads", alpsids)
	fmt.Println("After success launchpads", slpids)
	fmt.Println("After failed launchpads", flpids)

	k.SetCreatedLaunchpads(ctx,
		types.Launchpads{CreatedLaunchpads: clpsids})

	alpsidss = append(alpsidss, alpsids...)
	fmt.Println("appended active launchpad ids", alpsidss)
	k.SetActiveLaunchpads(ctx,
		types.Launchpads{ActiveLaunchpads: alpsidss})

	slps = append(slps, slpids...)
	k.SetSuccessLaunchpads(ctx,
		types.Launchpads{SuccessLaunchpads: slps})

	flps = append(flps, flpids...)
	k.SetFailLaunchpads(ctx,
		types.Launchpads{FailLaunchpads: flps})

	return nil
}
