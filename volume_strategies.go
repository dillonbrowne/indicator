// Copyright (c) 2021 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/indicator

package indicator

// Money flow index strategy.
func MoneyFlowIndexStrategy(asset Asset) []Action {
	actions := make([]Action, len(asset.Date))

	moneyFlowIndex := DefaultMoneyFlowIndex(
		asset.High,
		asset.Low,
		asset.Closing,
		asset.Volume)

	for i := 0; i < len(actions); i++ {
		if moneyFlowIndex[i] >= 80 {
			actions[i] = SELL
		} else {
			actions[i] = BUY
		}
	}

	return actions
}

// Force index strategy function.
func ForceIndexStrategy(asset Asset) []Action {
	actions := make([]Action, len(asset.Date))

	forceIndex := DefaultForceIndex(asset.Closing, asset.Volume)

	for i := 0; i < len(actions); i++ {
		if forceIndex[i] > 0 {
			actions[i] = BUY
		} else if forceIndex[i] < 0 {
			actions[i] = SELL
		} else {
			actions[i] = HOLD
		}
	}

	return actions
}