package cfgreader

//LHS basic operation
type LHS interface {
	getLHSRuleNum() uint16
	getRulesNum() uint16
	getAllRHS() []string
	addRule() bool
	deleteBeforeHandle(uint16) bool
	deleteRHSRule(uint16) bool
	deleteAfterHandle(uint16) bool
}

//RHS information operation
type RHS interface {
	getLHSRuleNum() uint16
	getAllToken() []string
}

type TokenInfo struct {
	token       string
	nonTerminal bool
}

type LHSInfo struct {
	ruleNum     uint16
	rulesNum    uint16
	nonTerminal TokenInfo
}

type RHSInfo struct {
	tokens []TokenInfo
}
