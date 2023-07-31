package prover_server

import "gate-zkmerkle-proof/global"

func Handler() {
	prover := NewProver(global.Cfg)
	prover.Run(false)
}
