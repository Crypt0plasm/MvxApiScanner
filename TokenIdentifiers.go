package MvxApiScanner

import (
	p "Firefly-APD"
)

// Types
type ESDT string
type MvxAddress string
type SFT string

type BalanceSFT struct {
	Address MvxAddress `json:"address"`
	Balance string     `json:"balance"`
}
type BalanceNFT struct {
	Address MvxAddress `json:"address"`
	NFT     NFT
}
type NFT struct {
	Collection string
	Identifier string
}
type BalanceESDT struct {
	Address MvxAddress `json:"address"`
	Balance string     `json:"balance"`
}

type TrueBalanceESDT struct {
	AB  BalanceESDT
	KYC bool
}

//VestaStructures

type VestaPool struct {
	VEGLD *p.Decimal
	Token *p.Decimal
}

type VestaSplit struct {
	Pool  VestaPool
	Vesta *p.Decimal
}

// Can be made with the following request url
// "https://api.multiversx.com/accounts/erd1h6lh2tqjscs4n69c4w4wunu4qw2mz708qn8mqk4quzsyz2syn0aq5gu64s/tokens/WEGLD-bd4d79"
// Then paste here: https://mholt.github.io/json-to-go/

type ESDTSuperStructure struct {
	Type          string `json:"type"`
	Identifier    string `json:"identifier"`
	Name          string `json:"name"`
	Ticker        string `json:"ticker"`
	Owner         string `json:"owner"`
	Minted        string `json:"minted"`
	Burnt         string `json:"burnt"`
	InitialMinted string `json:"initialMinted"`
	Decimals      int    `json:"decimals"`
	IsPaused      bool   `json:"isPaused"`
	Assets        struct {
		Website         string `json:"website"`
		Description     string `json:"description"`
		Status          string `json:"status"`
		PngURL          string `json:"pngUrl"`
		SvgURL          string `json:"svgUrl"`
		LedgerSignature string `json:"ledgerSignature"`
		Social          struct {
			Email      string `json:"email"`
			Twitter    string `json:"twitter"`
			Whitepaper string `json:"whitepaper"`
			Coingecko  string `json:"coingecko"`
			Discord    string `json:"discord"`
			Telegram   string `json:"telegram"`
		} `json:"social"`
		LockedAccounts         string   `json:"lockedAccounts"`
		ExtraTokens            []string `json:"extraTokens"`
		PreferredRankAlgorithm string   `json:"preferredRankAlgorithm"`
	} `json:"assets"`
	Transactions             int     `json:"transactions"`
	Accounts                 int     `json:"accounts"`
	CanUpgrade               bool    `json:"canUpgrade"`
	CanMint                  bool    `json:"canMint"`
	CanBurn                  bool    `json:"canBurn"`
	CanChangeOwner           bool    `json:"canChangeOwner"`
	CanPause                 bool    `json:"canPause"`
	CanFreeze                bool    `json:"canFreeze"`
	CanWipe                  bool    `json:"canWipe"`
	CanTransferNftCreateRole bool    `json:"canTransferNftCreateRole"`
	Price                    float64 `json:"price"`
	MarketCap                float64 `json:"marketCap"`
	Supply                   string  `json:"supply"`
	CirculatingSupply        string  `json:"circulatingSupply"`
	Timestamp                int     `json:"timestamp"`
	Balance                  string  `json:"balance"`
	ValueUsd                 float64 `json:"valueUsd"`
	Attributes               string  `json:"attributes"`
}

var (
	//TokenIdentifiers
	//ESDT Tokens
	WrappedEGLD = ESDT("WEGLD-bd4d79")
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
	CD10Binar    = SFT("DHCD-bc9963-0a")
	VestaGold    = SFT("VESTAXDAO-e6c48c-01")
	VestaSilver  = SFT("VESTAXDAO-e6c48c-02")
	VestaBronze  = SFT("VESTAXDAO-e6c48c-03")

	//NFTs
	SnakeNFT = SFT("DEMIOU-6d1b5c")

	//Liquidity Pools
	EGLDSuperLP = MvxAddress("erd1qqqqqqqqqqqqqpgqdx6z3sauy49c5k6c6lwhjqclrfwlxlud2jpsvwj5dp")
	EGLDCrustLP = MvxAddress("erd1qqqqqqqqqqqqqpgqj6daemefdk5kjgy9rs4zsng03kezgxdm2jps3h5n07")
	EGLDAeroLP  = MvxAddress("erd1qqqqqqqqqqqqqpgqzjctu8xrgn8jmfp503tajjvzz2zq60v92jpsslkh5a")

	//Smart Contracts
	//MintSC
	CodingDivisionMintSC = MvxAddress("erd1qqqqqqqqqqqqqpgqk7t2cc8awcgwnftnn4p9w7m8fjxplkfcj9qselntcv")
	VestaMinter          = MvxAddress("erd1qqqqqqqqqqqqqpgqtwe67hk3rwpjx2rky74csj069gftw32j2d2ssd2mvf")

	//MarketsSC
	MarketXoxno    = MvxAddress("erd1qqqqqqqqqqqqqpgq6wegs2xkypfpync8mn2sa5cmpqjlvrhwz5nqgepyg8")
	MarketFrameIt1 = MvxAddress("erd1qqqqqqqqqqqqqpgqx00c5udg9uql5sgk5vswfr8cp7kalqgcyawq9xpw26")
	MarketFrameIt2 = MvxAddress("erd1qqqqqqqqqqqqqpgq705fxpfrjne0tl3ece0rrspykq88mynn4kxs2cg43s")
	MarketNFTr     = MvxAddress("erd1qqqqqqqqqqqqqpgqz2tgn80j5p5hqh4hx69uc27gz0drcjmmg20skf0wru")
	Krogan         = MvxAddress("erd1qqqqqqqqqqqqqpgq8xwzu82v8ex3h4ayl5lsvxqxnhecpwyvwe0sf2qj4e")

	//SupercietyVaults SC
	SnakeDAO          = MvxAddress("erd1qqqqqqqqqqqqqpgql9zuu0r5pj9xcx94y08wujmwkn2hcuns27rsdcmzf0")
	CodingDivisionDAO = MvxAddress("erd1qqqqqqqqqqqqqpgqqkyp6auk2h7y6sdj2w55qp8zad5ddurn27rs47vs2n")
	VestaDAO          = MvxAddress("erd1qqqqqqqqqqqqqpgq9ez0cfmq7l9t7s800mdfagkekqj3g9dv27rsqjagnv")
	DHV1              = MvxAddress("erd1qqqqqqqqqqqqqpgqkpcm6xqrsfr8gzlh7pdlc2d6gfvfe2j827rstr8jwd")
	DHV2              = MvxAddress("erd1qqqqqqqqqqqqqpgq2k6pww09cegpgqpj5yctrx4tsxlry99g27rslvxyxm")
	DHV3              = MvxAddress("erd1qqqqqqqqqqqqqpgqrcd2dx3h5zpv6ngf7qgej95dkhejwv7m27rsxvf78a")

	//Addresses
	KosonicTreasury = MvxAddress("erd1h0ymqdgl6vf0pud0klz5nstwra3sxj06afaj86x0pg7p52dvve9qqtg7x4")
	Hefe            = MvxAddress("erd1vj40fxw0yah34mmdxly7l28w097ju6hf8pczpcdxs05n2vyx8hcspyxm2c")
)
