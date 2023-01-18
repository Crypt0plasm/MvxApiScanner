package MvxApiScanner

// MakeESDTSnapshotLink To Snapshot All Addresses that have a certain ESDT
func MakeESDTSnapshotLink(Address MvxAddress, Token ESDT) string {
	String1 := "https://api.multiversx.com/accounts/"
	String2 := "/tokens/"
	return String1 + string(Address) + String2 + string(Token)
}

func MakeAddressESDTSnapshotLink(Token ESDT) string {
	String1 := "https://api.multiversx.com/tokens/"
	String2 := "/accounts?size=10000"
	return String1 + string(Token) + String2
}

// MakeSFTSnapshotLink To Snapshot All Addresses that have a certain SFT
func MakeSFTSnapshotLink(Token SFT) string {
	String1 := "https://api.multiversx.com/nfts/"
	String2 := "/accounts?size=10000"
	return String1 + string(Token) + String2
}

// MakeNFTSnapshotLink To Snapshot the Address that has a certain NFT
func MakeNFTSnapshotLink(Collection SFT, Identifier string) string {
	String1 := "https://api.multiversx.com/nfts/"
	String2 := "/accounts"
	NFTDesignation := string(Collection) + "-" + Identifier

	return String1 + NFTDesignation + String2
}
