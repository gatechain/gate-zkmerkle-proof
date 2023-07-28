package config

import (
	"gate-zkmerkle-proof/utils"
	"math/big"
)

type Config struct {
	MysqlDataSource string
	UserDataFile    string
	DbSuffix        string
	TreeDB          struct {
		Driver string
		Option struct {
			Addr string
		}
	}
	Redis struct {
		Host     string
		Type     string
		Password string
	}
	ZkKeyName string
}

type CexConfig struct {
	ProofTable    string
	ZkKeyName     string
	CexAssetsInfo []utils.CexAssetInfo
}

type UserConfig struct {
	AccountIndex  uint32
	AccountIdHash string
	TotalEquity   big.Int
	TotalDebt     big.Int
	Root          string
	Assets        []utils.AccountAsset
	Proof         []string
}
