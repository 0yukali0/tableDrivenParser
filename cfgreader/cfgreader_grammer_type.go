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

//TokenInfo describe property of token
type TokenInfo struct {
	token       string
	nonTerminal bool
}

//LHSInfo describe arg of LHS
type LHSInfo struct {
	ruleNum     uint16
	rulesNum    uint16
	nonTerminal TokenInfo
}

//RHSInfo describe context of RHS
type RHSInfo struct {
	tokens []TokenInfo
}
