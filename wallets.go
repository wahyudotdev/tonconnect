package tonconnect

import (
	"slices"
)

type Wallet struct {
	Name         string `json:"name"`
	UniversalURL string `json:"universal_url"`
	BridgeURL    string `json:"bridge_url"`
}

var Wallets = map[string]Wallet{
	"telegram-wallet": {
		Name:         "Wallet",
		UniversalURL: "https://t.me/wallet/start?startapp=",
		BridgeURL:    "https://bridge.ton.space/bridge",
	},
	"tonkeeper": {
		Name:         "Tonkeeper",
		UniversalURL: "https://app.tonkeeper.com/ton-connect",
		BridgeURL:    "https://bridge.tonapi.io/bridge",
	},
	"mytonwallet": {
		Name:         "MyTonWallet",
		UniversalURL: "https://connect.mytonwallet.org/",
		BridgeURL:    "https://tonconnectbridge.mytonwallet.org/bridge",
	},
	"tonhub": {
		Name:         "Tonhub",
		UniversalURL: "https://tonhub.com/ton-connect",
		BridgeURL:    "https://connect.tonhubapi.com/tonconnect",
	},
	"dewallet": {
		Name:         "DeWallet",
		UniversalURL: "https://t.me/dewallet?attach=wallet",
		BridgeURL:    "https://sse-bridge.delab.team/bridge",
	},
	"bitgetTonWallet": {
		Name:         "Bitget Wallet",
		UniversalURL: "https://bkcode.vip/ton-connect",
		BridgeURL:    "https://bridge.tonapi.io/bridge",
	},
	"safepalwallet": {
		Name:         "SafePal",
		UniversalURL: "https://link.safepal.io/ton-connect",
		BridgeURL:    "https://ton-bridge.safepal.com/tonbridge/v1/bridge",
	},
}

func getBridgeURLs(wallets ...Wallet) []string {
	var bridges []string
	for _, w := range wallets {
		bridges = append(bridges, w.BridgeURL)
	}

	slices.Sort(bridges)
	return slices.Compact(bridges)
}
