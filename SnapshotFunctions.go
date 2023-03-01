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
// 	[C]01         MakeNFTBaseString				Makes the Base NFT String
// 	[C]02         Make20DigitsNFTString			Makes the 20 Digit Base NFT String
// 	[C]03         Make256NFTString				Makes the 20 Digit Base NFT String
// 	[C]04         MakeTotalNFTString			Makes the Final NFT Identifiers in a slice of String
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
//	Makes the Base NFT String

func MakeNFTBaseString() []string {
	var (
		SingleString string
		StringSlice  []string
	)
	for i := int64(0); i < 16; i++ {
		SingleString = strconv.FormatInt(i, 16)
		StringSlice = append(StringSlice, SingleString)
	}
	return StringSlice
}

// Make20DigitsNFTString ===============================================================================================
//
// [C]02         Make20DigitsNFTString
//
//	Makes the 20 Digit Base NFT String
func Make20DigitsNFTString(Number int64) []string {
	var (
		String       string
		OutputString []string
	)
	BS := MakeNFTBaseString()

	if Number >= 0 && Number < 16 {
		SS := strconv.FormatInt(Number, 16)
		for i := 0; i < len(BS); i++ {
			String = SS + BS[i]
			OutputString = append(OutputString, String)
		}
	} else {
		return BS
	}
	return OutputString
}

// Make256NFTString ====================================================================================================
//
// [C]03         Make256NFTString
//
//	Makes the 20 Digit Base NFT String
func Make256NFTString() []string {
	var (
		SingleStringChain []string
		String256         []string
	)
	for i := int64(0); i < 16; i++ {
		SingleStringChain = Make20DigitsNFTString(i)
		String256 = append(String256, SingleStringChain...)
	}
	return String256
}

// MakeTotalNFTString ====================================================================================================
//
// [C]04         MakeTotalNFTString
//
//	Makes the Final NFT Identifiers in a slice of String
func MakeTotalNFTString(Size int64) []string {
	var (
		Output   []string
		Addition string
		ToAppend []string
	)

	AppendedStringChain := func(Prefix string, StringChain []string) []string {
		var (
			Unit   string
			Result []string
		)
		for i := 0; i < len(StringChain); i++ {
			Unit = Prefix + StringChain[i]
			Result = append(Result, Unit)
		}
		return Result
	}
	GetPrefix := func(Number int) string {
		var Prefix string
		if Number < 10 {
			Prefix = "0" + strconv.Itoa(Number)
		} else {
			Prefix = strconv.Itoa(Number)
		}
		return Prefix
	}

	Whole := mt.DivInt(p.NFI(Size), p.NFS("256"))
	Rest := mt.DivMod(p.NFI(Size), p.NFS("256"))
	WholeInt, _ := strconv.Atoi(mt.DTS(Whole))
	RestInt, _ := strconv.Atoi(mt.DTS(Rest))

	String256 := Make256NFTString()

	if WholeInt == 0 {
		Output = String256[:RestInt+1]
	} else {
		for i := 0; i < WholeInt+1; i++ {
			if i == 0 {
				Output = append(Output, String256...)
			} else if i >= 1 && i < WholeInt {
				Addition = GetPrefix(i)
				ToAppend = AppendedStringChain(Addition, String256)
				Output = append(Output, ToAppend...)
			} else if i == WholeInt {
				Addition = GetPrefix(i)
				ToAppend = AppendedStringChain(Addition, String256[:RestInt+1])
				Output = append(Output, ToAppend...)
			}
		}
	}
	return Output
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
