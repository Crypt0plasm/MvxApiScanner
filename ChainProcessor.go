package MvxApiScanner

import (
	p "Firefly-APD"
	mt "SuperMath"
)

//
// MATH
//

// AddBalanceIntegerChain =============================================================================================
//
// [A]01         AddBalanceSFTChain
//
//	Add the Balances of an SFT Chain
func AddBalanceIntegerChain(Chain []BalanceSFT) *p.Decimal {
	Sum := p.NFS("0")
	for i := 0; i < len(Chain); i++ {
		Sum = mt.ADDxc(Sum, p.NFS(Chain[i].Balance))
	}
	return Sum
}
func AddBalanceDecimalChain(Chain []BalanceESDT) *p.Decimal {
	Sum := p.NFS("0")
	for i := 0; i < len(Chain); i++ {
		Sum = mt.ADDxc(Sum, p.NFS(Chain[i].Balance))
	}
	return Sum
}

// RewardsComputerIntegerChain =============================================================================================
//
// [A]02         RewardsComputerIntegerChain
//
//		Computes Reward for an Integer Chain of Values
//		Reward string, can be a decimal with up to 18 decimals.
//	 Output must be BalanceESDT due to decimals
func RewardsComputerIntegerChain(Input []BalanceSFT, Reward string) []BalanceESDT {
	var (
		Unit   BalanceESDT
		Output []BalanceESDT
	)
	for i := 0; i < len(Input); i++ {
		Unit.Address = Input[i].Address
		Unit.Balance = mt.DTS(mt.MULxc(p.NFS(Input[i].Balance), p.NFS(Reward)))
		Output = append(Output, Unit)
	}
	return Output
}

// DecimalChainAdder =============================================================================================
//
// [A]03         DecimalChainAdder
//
// Adds two BalanceESDT Files, removing duplicate addresses and summing their balances.
func DecimalChainAdder(S1, S2 []BalanceESDT) []BalanceESDT {
	var (
		ValueS1, ValueS2, TotalValue *p.Decimal
	)
	AllSlice := append(S1, S2...)

	//2.    Make a slice with all Elrond Address (will contain duplicate Elrond Addresses)
	//      basically removes the balance value.
	ElrondSlice := make([]MvxAddress, len(AllSlice))

	for i := 0; i < len(AllSlice); i++ {
		ElrondSlice[i] = AllSlice[i].Address
	}

	Unique := RemoveDuplicateMvxAddresses(ElrondSlice)
	Output := make([]BalanceESDT, len(Unique))

	for i := 0; i < len(Unique); i++ {
		for j := 0; j < len(S1); j++ {
			if Unique[i] == S1[j].Address {
				ValueS1 = p.NFS(S1[j].Balance)
				break
			} else {
				ValueS1 = p.NFS("0")
			}
		}

		for k := 0; k < len(S2); k++ {
			if Unique[i] == S2[k].Address {
				ValueS2 = p.NFS(S2[k].Balance)
				break
			} else {
				ValueS2 = p.NFS("0")
			}
		}

		TotalValue = mt.ADDxc(ValueS1, ValueS2)

		Output[i].Address = Unique[i]
		Output[i].Balance = mt.DTS(TotalValue)
	}
	return Output
}

//
// Chain Processing
//

// RemoveDuplicateMvxAddresses =========================================================================================
//
// [B]01         func RemoveDuplicateMvxAddresses(Input []ElrondAddress) []ElrondAddress {
//
//	Add the Balances of an SFT Chain
func RemoveDuplicateMvxAddresses(Input []MvxAddress) []MvxAddress {
	//3.    Remove Duplicate Elrond Addresses.
	//3.1   Make a hash map from ElrondAddress to int
	Check := make(map[MvxAddress]int)

	//3.2   Make the empty Output Slice that will contain unique Elrond Addresses
	Unique := make([]MvxAddress, 0)

	//3.3   Iterate through the Slice containing duplicates and map each element to 0. (or any other thing)
	for _, val := range Input {
		Check[val] = 1
	}

	//3.4   Now finally iterate through the map and append each key of the map to a
	//      new slice of strings. Since any duplicate value too will be mapped to the
	//      same number as the previous one, hence all the keys will be unique.
	for letter, _ := range Check {
		Unique = append(Unique, letter)
	}
	return Unique
}

// SortBalanceIntegerChain =============================================================================================
//
// [B]02         SortBalanceIntegerChain
//
//	Sorts a BalanceSFTChain Chain from the highest Balance to the lowest Balance
func SortBalanceIntegerChain(Chain []BalanceSFT) []BalanceSFT {
	var (
		SortedChain []BalanceSFT
	)
	GetMaxElement := func(Chain []BalanceSFT) int {
		Max := 0
		for i := 0; i < len(Chain); i++ {
			if Chain[i].Balance >= Chain[Max].Balance == true {
				Max = i
			}
		}
		return Max
	}

	Chain2Sort := Chain

	for i := 0; i < len(Chain); i++ {
		Biggest := GetMaxElement(Chain2Sort)
		Unit := BalanceSFT{Address: Chain2Sort[Biggest].Address, Balance: Chain2Sort[Biggest].Balance}
		SortedChain = append(SortedChain, Unit)

		//Removing the biggest element
		//This syntax removes from a slice the element on position Biggest
		Chain2Sort = append(Chain2Sort[:Biggest], Chain2Sort[Biggest+1:]...)
	}
	return SortedChain
}

// SortBalanceDecimalChain =============================================================================================
//
// [B]03         SortBalanceDecimalChain
//
//	Sorts a BalanceSFTChain Chain from the highest Balance to the lowest Balance
func SortBalanceDecimalChain(Chain []BalanceESDT) []BalanceESDT {
	var (
		SortedChain []BalanceESDT
	)
	GetMaxElement := func(Chain []BalanceESDT) int {
		Max := 0
		for i := 0; i < len(Chain); i++ {
			if mt.DecimalGreaterThanOrEqual(p.NFS(Chain[i].Balance), p.NFS(Chain[Max].Balance)) == true {
				Max = i
			}
		}
		return Max
	}
	Chain2Sort := Chain

	for i := 0; i < len(Chain); i++ {
		Biggest := GetMaxElement(Chain2Sort)
		Unit := BalanceESDT{Address: Chain2Sort[Biggest].Address, Balance: Chain2Sort[Biggest].Balance}
		SortedChain = append(SortedChain, Unit)

		//Removing the biggest element
		//This syntax removes from a slice the element on position Biggest
		Chain2Sort = append(Chain2Sort[:Biggest], Chain2Sort[Biggest+1:]...)
	}
	return SortedChain
}

// ConvertSFTtoESDTChain =============================================================================================
//
// [C]01         ConvertSFTtoESDTChain
//
//		Sorts a BalanceSFTChain Chain from the highest Balance to the lowest Balance
//	 Converts Integer balances that are Atomic Unics in Decimals
func ConvertSFTtoESDTChain(Input []BalanceSFT) []BalanceESDT {
	Output := make([]BalanceESDT, len(Input))

	for i := 0; i < len(Input); i++ {
		Output[i] = ConvertSFTtoESDT(Input[i])
	}
	return Output
}

func ConvertSFTtoESDT(Input BalanceSFT) (Output BalanceESDT) {
	Output.Address = Input.Address
	Output.Balance = mt.DTS(mt.TruncateCustom(mt.DIVxc(p.NFS(Input.Balance), p.NFS("1000000000000000000")), 18))
	return
}
