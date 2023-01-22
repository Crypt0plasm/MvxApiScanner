package MvxApiScanner

import (
	p "Firefly-APD"
	mt "SuperMath"
	"fmt"
	"os"
)

//
//
//
// MATH Functions
//
// [A]01		AddBalanceDecimalChain							Decimal version (ESDT)
// [A]01b		AddBalanceIntegerChain							Integer Version (SFT)
// [A]02		RewardsComputerDecimalChainMultiplication		Decimal Version (ESDT)
// [A]02b		RewardsComputerIntegerChainMultiplication		Integer Version (SFT)
// [A]03		DecimalChainAdder								Decimal Version (ESDT)
// [A]03b		IntegerChainAdder								Integer Version (SFT)
// [A]04		ExactPercentualDecimalRewardSplitter			Decimal Version (ESDT)
// [A]04b		ExactPercentualIntegerRewardSplitter			Integer Version (SFT)
//
// CHAIN PROCESSING
//
// [B]01		RemoveDuplicateMvxAddresses
// [B]02		SortBalanceDecimalChain							Decimal Version (ESDT)
// [B]02b		SortBalanceIntegerChain							Integer Version (SFT)
//
// SFT to ESDT Convertors
//
// [C]01		ConvertSFTAUStoESDTChain						Converts a chain of large integers (AUs) to ESDT Decimal
// [C]01b		ConvertAUStoESDT								Converts a large integer (AUS) to ESDT Decimal
// [C]02		ConvertIntegerSFTtoESDTChain					Converts a SFT Chain to ESDT Chain
// [C]02b		ConvertIntegerSFTtoESDT							Converts a SFT to ESDT
//
//
//

// AddBalanceDecimalChain =============================================================================================
//
// [A]01         AddBalanceESDTChain
//
//	Add the Balances of an ESDT Chain
func AddBalanceDecimalChain(Chain []BalanceESDT) *p.Decimal {
	Sum := p.NFS("0")
	for i := 0; i < len(Chain); i++ {
		Sum = mt.ADDxc(Sum, p.NFS(Chain[i].Balance))
	}
	return Sum
}

// AddBalanceIntegerChain
// [A]01b         AddBalanceSFTChain
// Integer Version (SFT Input)
func AddBalanceIntegerChain(Chain []BalanceSFT) *p.Decimal {
	return AddBalanceDecimalChain(ConvertIntegerSFTtoESDTChain(Chain))
}

// RewardsComputerDecimalChainMultiplication ===========================================================================
//
// [A]02         RewardsComputerIntegerChain
//
//		Computes Reward for a Decimal Chain of Values
//		Reward string, can be a decimal with up to 18 decimals.
//	 Output must be BalanceESDT due to decimals
//	 Multiplies Balances with Reward-Value; Simple multiplication.
func RewardsComputerDecimalChainMultiplication(Input []BalanceESDT, Reward string) []BalanceESDT {
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

// RewardsComputerIntegerChainMultiplication Function with Integers (SFT) as input
// [A]02b         RewardsComputerIntegerChain
// Integer Version (SFT Input)
func RewardsComputerIntegerChainMultiplication(Input []BalanceSFT, Reward string) []BalanceESDT {
	return RewardsComputerDecimalChainMultiplication(ConvertIntegerSFTtoESDTChain(Input), Reward)
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

// IntegerChainAdder Function with Integers (SFT) as input
// [A]03b
// Integer Version (SFT Input)
func IntegerChainAdder(S1, S2 []BalanceSFT) []BalanceESDT {
	S1ESDT := ConvertIntegerSFTtoESDTChain(S1)
	S2ESDT := ConvertIntegerSFTtoESDTChain(S2)
	return DecimalChainAdder(S1ESDT, S2ESDT)
}

// PercentualSplitter ================================================================================
//
// [A]04
// The Base Function that computes a percentual split based on a given total value and a slice of Decimals
func PercentualSplitter(TotalSplit *p.Decimal, Input []*p.Decimal) []*p.Decimal {
	Output := make([]*p.Decimal, len(Input))
	Sum := p.NFS("0")

	var (
		Remainder       *p.Decimal
		Last            *p.Decimal
		IndividualSplit *p.Decimal
	)

	for i := 0; i < len(Input); i++ {
		Sum = mt.ADDxc(Sum, Input[i])
	}

	for i := 0; i < len(Input); i++ {
		IndividualSplit = mt.TruncateCustom(mt.MULxc(mt.DIVxc(Input[i], Sum), TotalSplit), 18)
		if i == 0 {
			Remainder = mt.SUBxc(TotalSplit, IndividualSplit)
		} else {
			Remainder = mt.SUBxc(Remainder, IndividualSplit)
		}
		if i == len(Input)-2 {
			Last = Remainder
		}

		if i == len(Input)-1 {
			Output[i] = Last
		} else {
			Output[i] = IndividualSplit
		}
	}
	return Output
}

// ExactPercentualDecimalRewardSplitter ================================================================================
//
// [A]04
// Computes IndividualSplits, using a given amount of Token, for a chain of ESDT Balances
// Resulted Chain of ESDT Balances represents amounts based on percents
//
//	from the total Balance Amount based of Input ESDT Balances.
func ExactPercentualDecimalRewardSplitter(AmountToSplit *p.Decimal, InputChain []BalanceESDT) []BalanceESDT {
	InputDecimalSlice := make([]*p.Decimal, len(InputChain))
	for i := 0; i < len(InputChain); i++ {
		InputDecimalSlice[i] = p.NFS(InputChain[i].Balance)
	}
	OutputDecimalSlice := PercentualSplitter(AmountToSplit, InputDecimalSlice)

	Output := make([]BalanceESDT, len(InputChain))
	for i := 0; i < len(InputChain); i++ {
		Output[i].Address = InputChain[i].Address
		Output[i].Balance = mt.DTS(OutputDecimalSlice[i])
	}
	return Output
}

// ExactPercentualIntegerRewardSplitter Function with Integers (SFT) as input
// [A]04b
// Integer Version (SFT Input)
func ExactPercentualIntegerRewardSplitter(AmountToSplit *p.Decimal, InputChain []BalanceSFT) []BalanceESDT {
	return ExactPercentualDecimalRewardSplitter(AmountToSplit, ConvertIntegerSFTtoESDTChain(InputChain))
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

// SortBalanceDecimalChain =============================================================================================
//
// [B]02         SortBalanceDecimalChain
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

// SortBalanceIntegerChain Function with Integers (SFT) as input
// [B]02b
// Integer Version (SFT Input)
func SortBalanceIntegerChain(Chain []BalanceSFT) []BalanceESDT {
	return SortBalanceDecimalChain(ConvertIntegerSFTtoESDTChain(Chain))
}

// ConvertSFTAUStoESDTChain ============================================================================================
//
// [C]01
//
//	Converts Integer balances that are Atomic Unics in Decimals
func ConvertSFTAUStoESDTChain(Input []BalanceSFT) []BalanceESDT {
	Output := make([]BalanceESDT, len(Input))

	for i := 0; i < len(Input); i++ {
		Output[i] = ConvertAUStoESDT(Input[i])
	}
	return Output
}

// ConvertAUStoESDT
// [C]01b
// Individual Function
func ConvertAUStoESDT(Input BalanceSFT) (Output BalanceESDT) {
	Output.Address = Input.Address
	Output.Balance = mt.DTS(mt.TruncateCustom(mt.DIVxc(p.NFS(Input.Balance), p.NFS("1000000000000000000")), 18))
	return
}

// ConvertIntegerSFTtoESDTChain ============================================================================================
//
// [C]02
// Converts Integer chain to Decimal Chain (Type conversion from BalanceSFT to BalanceESDT
func ConvertIntegerSFTtoESDTChain(Input []BalanceSFT) []BalanceESDT {
	Output := make([]BalanceESDT, len(Input))

	for i := 0; i < len(Input); i++ {
		Output[i] = ConvertIntegerSFToESDT(Input[i])
	}
	return Output
}

// ConvertIntegerSFToESDT
// [C]02b
// Individual Function
func ConvertIntegerSFToESDT(Input BalanceSFT) (Output BalanceESDT) {
	Output.Address = Input.Address
	Output.Balance = Input.Balance
	return
}

func ConvertToBulkCSV(OutputName string, InputChain []BalanceESDT) {
	f, err := os.Create(OutputName)
	if err != nil {
		fmt.Println(err)
		_ = f.Close()
		return
	}

	LineToPrint := func(Unit BalanceESDT) string {
		L := string(Unit.Address) + ";" + Unit.Balance
		return L
	}

	for i := 0; i < len(InputChain); i++ {
		_, err1 := fmt.Fprintln(f, LineToPrint(InputChain[i]))
		if err1 != nil {
			return
		}
	}
}
