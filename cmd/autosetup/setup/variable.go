package setup

const (
	DefaultValueSignal = "d"
	EmptyLine          = ""
	Positive           = "yes"
	Negative           = "no"
	Separator          = ","
)

// read type
const (
	ClearData  = "ClearData"
	IsBp     = "IsBp"
	YesOrNo  = "YesOrNo"
	NodeName = "NodeName"
	ChainId  = "ChainId"
	BpName   = "BpName"
	PriKey   = "PriKey"
	SeedList = "SeedList"
	LogLevel = "LogLevel"
	DataDir  = "DataDir"

	StartNode = "StartNode"
)

var (
	InitNewConfig = true
	NodeIsBp      = false
	StartNodeNow  = false
	ClearLocalData = false
)