package MvxApiScanner

import "encoding/json"

// ======================================================================================================================
//
//	SnapshotFunctions.go
//	Blockchain Snapshot Functions
//
// 	[A]01         SnapshotIntegerChain			Creates a Chain of integer values (SFT Values)
// 	[A]01         SnapshotDecimalChain			Creates a Chain of Decimal values (ESDT Values)

func SnapshotIntegerChain(Link string) []BalanceSFT {
	var OutputChain []BalanceSFT
	SS := OnPage(Link)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

func SnapshotDecimalChain(Link string) []BalanceESDT {
	var OutputChain []BalanceESDT
	SS := OnPage(Link)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}
