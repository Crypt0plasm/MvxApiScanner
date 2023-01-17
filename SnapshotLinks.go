package MvxApiScanner

func MakeESDTSnapshotLink(Token ESDT) string {
	String1 := "https://api.multiversx.com/tokens/"
	String2 := "/accounts?size=10000"
	return String1 + string(Token) + String2
}

func MakeSFTSnapshotLink(Token SFT) string {
	String1 := "https://api.multiversx.com/nft/"
	String2 := "/accounts?size=10000"
	return String1 + string(Token) + String2
}

func MakeSNFTSnapshotLink(Token SFT) string {
	String1 := "https://api.multiversx.com/nft/"
	String2 := "/accounts?size=1"
	return String1 + string(Token) + String2
}
