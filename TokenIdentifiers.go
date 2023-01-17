package MvxApiScanner

import (
	p "github.com/Crypt0plasm/Firefly-APD"
)

// Types
type ESDT string
type MvxAddress string
type SFT string
type NFT string
type BalanceSFT struct {
	Address MvxAddress
	Balance int
}
type BalanceESDT struct {
	Address MvxAddress
	Balance *p.Decimal
}

var (
	//TokenIdentifiers
	//ESDT Tokens
	wrappedEGLD = ESDT("WEGLD-bd4d79")
	vestaEGLD   = ESDT("VEGLD-2b9319")

	Super = ESDT("SUPER-507aa6")
	XLH   = ESDT("XLH-8daa50")
	Crust = ESDT("CRU-a5f4aa")
	Aero  = ESDT("AERO-458bbf")

	//LP Tokens
	SUPEREGLD = ESDT("SUPEREGLD-a793b9")
	CRUSTEGLD = ESDT("CRUWEGLD-76c269")
	AEROEGLD  = ESDT("AEROWEGLD-81cc37")

	//SFTs
	CD01SnakeEye = SFT("DHCD-bc9963-01")
	CD02Rudis    = SFT("DHCD-bc9963-02")
	CD03Gwen     = SFT("DHCD-bc9963-03")
	CD04Clutter  = SFT("DHCD-bc9963-04")
	CD05Bangai   = SFT("DHCD-bc9963-05")
	CD06Binos    = SFT("DHCD-bc9963-06")
	CD07Rubia    = SFT("DHCD-bc9963-07")
	CD08Ocultus  = SFT("DHCD-bc9963-08")
	CD09Oreta    = SFT("DHCD-bc9963-09")
	CD10Binar    = SFT("DHCD-bc9963-10")
	VestaGold    = SFT("VESTAXDAO-e6c48c-01")

	//NFTs
	SnakeNFT = NFT("DEMIOU-6d1b5c")

	//Liquidity Pools
	EGLDSuperLP = MvxAddress("erd1qqqqqqqqqqqqqpgqdx6z3sauy49c5k6c6lwhjqclrfwlxlud2jpsvwj5dp")
	EGLDCrustLP = MvxAddress("erd1qqqqqqqqqqqqqpgqj6daemefdk5kjgy9rs4zsng03kezgxdm2jps3h5n07")
	EGLDAeroLP  = MvxAddress("erd1qqqqqqqqqqqqqpgqzjctu8xrgn8jmfp503tajjvzz2zq60v92jpsslkh5a")

	//Smart Contracts
	//MintSC
	CodingDivisionMintSC = MvxAddress("erd1qqqqqqqqqqqqqpgqk7t2cc8awcgwnftnn4p9w7m8fjxplkfcj9qselntcv")

	//MarketsSC
	MarketXoxno    = MvxAddress("erd1qqqqqqqqqqqqqpgq6wegs2xkypfpync8mn2sa5cmpqjlvrhwz5nqgepyg8")
	MarketFrameIt1 = MvxAddress("erd1qqqqqqqqqqqqqpgqx00c5udg9uql5sgk5vswfr8cp7kalqgcyawq9xpw26")
	MarketFrameIt2 = MvxAddress("erd1qqqqqqqqqqqqqpgq705fxpfrjne0tl3ece0rrspykq88mynn4kxs2cg43s")
	MarketNFTr     = MvxAddress("erd1qqqqqqqqqqqqqpgqz2tgn80j5p5hqh4hx69uc27gz0drcjmmg20skf0wru")
	Krogan         = MvxAddress("erd1qqqqqqqqqqqqqpgq8xwzu82v8ex3h4ayl5lsvxqxnhecpwyvwe0sf2qj4e")

	//SupercietyVaults SC
	SupercietyVaultSnake = MvxAddress("erd1qqqqqqqqqqqqqpgql9zuu0r5pj9xcx94y08wujmwkn2hcuns27rsdcmzf0")
	SupercietyVaultCD    = MvxAddress("erd1qqqqqqqqqqqqqpgqqkyp6auk2h7y6sdj2w55qp8zad5ddurn27rs47vs2n")
	SupercietyVaultVesta = MvxAddress("erd1qqqqqqqqqqqqqpgq9ez0cfmq7l9t7s800mdfagkekqj3g9dv27rsqjagnv")

	//Addresses
	KosonicTreasury = MvxAddress("erd1h0ymqdgl6vf0pud0klz5nstwra3sxj06afaj86x0pg7p52dvve9qqtg7x4")
	VestaMinter     = MvxAddress("erd1qqqqqqqqqqqqqpgqtwe67hk3rwpjx2rky74csj069gftw32j2d2ssd2mvf")
)
