package openwtester

import (
	"github.com/Assetsadapter/dogecoin-adapter/dogecoin"
	//"github.com/blocktree/bitcoin-adapter/dogecoin"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(dogecoin.Symbol, dogecoin.NewWalletManager())
}
