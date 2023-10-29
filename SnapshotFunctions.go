package MvxApiScanner

import (
	p "Firefly-APD"
	mt "SuperMath"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// ======================================================================================================================
//
//	SnapshotFunctions.go
//	Blockchain Snapshot Functions
//
// 	[A]01         SnapshotIntegerChain			Creates a Chain of integer values (SFT Values)
//  [A]01b        SnapshotSFTChain				Snapshot SFT Chain
// 	[A]02         SnapshotDecimalChain			Creates a Chain of Decimal values (ESDT Values)
//
// 	[B]01         GetAddressIntegerAmount		Gets the Amount linked to a given address from an Integer Chain
// 	[B]01b        GetAddressDecimalAmount		Gets the Amount linked to a given address from an Integer Chain
// 	[B]02         GetAddressESDTAmount			Gets the Amount of ESDT Token from a given Address
//
// 	[C]01         MvxNftId						Makes the Base NFT String Identifier
//
//	[D]01		  ReadOmniscientFile
//
//
//

// SnapshotIntegerChain ================================================
//
// [A]01         SnapshotIntegerChain
//
// Snapshots a Link to an Integer Chain
func SnapshotIntegerChain(Link string) []BalanceSFT {
	var OutputChain []BalanceSFT
	SS := OnPage(Link)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

// SnapshotSFTChain ================================================
//
// [A]01b         SnapshotDecimalChain
//
// Snapshots a Link to a Decimal Chain
func SnapshotSFTChain(SemiFungibleToken SFT) []BalanceSFT {
	SFTLink := MakeSFTSnapshotLink(SemiFungibleToken)
	return SnapshotIntegerChain(SFTLink)
}

// SnapshotDecimalChain ================================================
//
// [A]02         SnapshotDecimalChain
//
// Snapshots a Link to a Decimal Chain
func SnapshotDecimalChain(Link string) []BalanceESDT {
	var (
		Unit        BalanceESDT
		OutputChain []BalanceESDT
	)

	Chain := SnapshotIntegerChain(Link)
	for i := 0; 1 < len(Chain); i++ {
		Unit.Address = Chain[i].Address
		Unit.Balance = mt.DTS(AtomicUnitsToPointDecimal(p.NFS(Chain[i].Balance)))
	}

	return OutputChain
}

// GetAddressIntegerAmount =============================================================================================
//
// [B]01         GetAddressAmount
//
//	Gets the Amount for a given address in an Integer Chain
func GetAddressIntegerAmount(Addy MvxAddress, Chain []BalanceSFT) string {
	var Result string
	for i := 0; i < len(Chain); i++ {
		if Chain[i].Address == Addy {
			Result = Chain[i].Balance
			break
		} else {
			Result = "0"
		}
	}
	return Result
}

// GetAddressDecimalAmount =============================================================================================
//
// [B]01b         GetAddressAmount
//
//	Gets the Amount for a given address in an Decimal Chain
func GetAddressDecimalAmount(Addy MvxAddress, Chain []BalanceESDT) string {
	var Result string
	for i := 0; i < len(Chain); i++ {
		if Chain[i].Address == Addy {
			Result = Chain[i].Balance
			break
		} else {
			Result = "0"
		}
	}
	return Result
}

// GetAddressESDTAmount =============================================================================================
//
// [B]02         GetAddressESDTAmount
//
//	Gets the Amount of ESDT Token from a given Address
func GetAddressESDTAmount(Addy MvxAddress, Token ESDT) *p.Decimal {
	var (
		ScannedJSON ESDTSuperStructure
		Balance     *p.Decimal
	)
	ScanURL := MakeESDTSnapshotLink(Addy, Token)

	Snapshot := OnPage(ScanURL)
	_ = json.Unmarshal([]byte(Snapshot), &ScannedJSON)
	if Snapshot == "[]" {
		Balance = p.NFS("0")
	} else {
		Balance = p.NFS(ScannedJSON.Balance)
	}
	return AtomicUnitsToPointDecimal(Balance)
}

// NFT Scan Functions
// MakeNFTBaseString ==================================================================================================
//
// [C]01         MakeNFTBaseString
//
//		Makes the Base NFT String
//	 Returns the ID as used by MVX NFTs
func MvxNftId(Number int64) string {
	var Output string
	RawString := strconv.FormatInt(Number, 16)
	if len(RawString) == 1 || len(RawString) == 3 {
		Output = "0" + RawString
	} else {
		Output = RawString
	}
	return Output
}

func MakeMVXNftIdSlice(Number int64) []string {
	var (
		String       string
		OutputString []string
	)

	for i := int64(1); i < Number; i++ {
		String = MvxNftId(int64(i))
		OutputString = append(OutputString, String)
	}
	return OutputString
}

//

func TripleDigitDesignation(Number int, Designation string) (StringName string) {
	if Number < 10 {
		StringName = Designation + "00" + strconv.Itoa(Number)
	} else if Number >= 10 && Number < 100 {
		StringName = Designation + "0" + strconv.Itoa(Number)
	} else {
		StringName = Designation + strconv.Itoa(Number)
	}
	return
}

func IzFile(Path, Filename string) bool {
	info, err := os.Stat(Path + Filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ReadOmniscientSS(Path, Filename string) []BalanceSFT {
	var (
		MainUnit BalanceSFT
		Output   []BalanceSFT
	)

	ProcessLine := func(Line string) BalanceSFT {
		var (
			Unit BalanceSFT
		)
		LineString := strings.ReplaceAll(Line, "\"", "")
		SeparatedLineStringSlice := strings.Split(LineString, ",")
		Unit.Address = MvxAddress(SeparatedLineStringSlice[0])
		Unit.Balance = SeparatedLineStringSlice[1]
		return Unit
	}

	StringSlice := ReadFile(Path + Filename)
	for i := 0; i < len(StringSlice); i++ {
		MainUnit = ProcessLine(StringSlice[i])
		Output = append(Output, MainUnit)
	}
	return Output
}

func ReadVestaPoolSnapshot(Path, Filename string) []VestaPool {
	var Output []VestaPool
	if IzFile(Path, Filename) == true {
		Output = ReadVestaPoolChain(Path, Filename)
	}
	return Output
}

func ReadBalanceChain(Path, Filename string) []BalanceESDT {
	var Output []BalanceESDT
	if IzFile(Path, Filename) == true {
		Output = ReadChain(Path, Filename)
	} else {
		Output = AdditionNeutralBalanceESDTChain
	}
	return Output
}

func ReadVestaPoolChain(Path, Filename string) []VestaPool {
	var (
		MainUnit VestaPool
		Output   []VestaPool
	)
	ProcessLine := func(Line string) VestaPool {
		var (
			Unit VestaPool
		)
		LineString := strings.ReplaceAll(Line, "{", "")
		LineString2 := strings.ReplaceAll(LineString, "}", "")
		SeparatedLineStringSlice := strings.Split(LineString2, " ")
		Unit.VEGLD = p.NFS(SeparatedLineStringSlice[0])
		Unit.Token = p.NFS(SeparatedLineStringSlice[1])
		return Unit
	}

	StringSlice := ReadFile(Path + Filename)
	for i := 0; i < len(StringSlice); i++ {
		MainUnit = ProcessLine(StringSlice[i])
		Output = append(Output, MainUnit)
	}
	return Output
}

func ReadChain(Path, Filename string) []BalanceESDT {
	var (
		MainUnit BalanceESDT
		Output   []BalanceESDT
	)

	ProcessLine := func(Line string) BalanceESDT {
		var (
			Unit BalanceESDT
		)
		LineString := strings.ReplaceAll(Line, "{", "")
		LineString2 := strings.ReplaceAll(LineString, "}", "")
		SeparatedLineStringSlice := strings.Split(LineString2, " ")
		Unit.Address = MvxAddress(SeparatedLineStringSlice[0])
		Unit.Balance = SeparatedLineStringSlice[1]
		return Unit
	}

	TotalPath := Path + Filename

	StringSlice := ReadFile(TotalPath)
	for i := 0; i < len(StringSlice); i++ {
		MainUnit = ProcessLine(StringSlice[i])
		Output = append(Output, MainUnit)
	}
	return Output
}
