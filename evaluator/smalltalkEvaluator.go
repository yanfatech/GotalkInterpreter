package evaluator

import (
	"GotalkInterpreter/parser"
	"GotalkInterpreter/treeNodes"
)

//testing stuff
func NewTestEvaluator() *Evaluator {
	evaluator := new(Evaluator)
	evaluator.globalScope = new(treeNodes.Scope).Initialize()
	return evaluator
}

func TestEval(codeString string) treeNodes.SmalltalkObjectInterface {
	evaluator := NewTestEvaluator()
	programNode := parser.InitializeParserFor(codeString)
	return evaluator.EvaluateProgramNode(programNode)
}

func TestEvalWithScope(codeString string, scope *treeNodes.Scope) treeNodes.SmalltalkObjectInterface {
	evaluator := NewEvaluatorWithGlobalScope(scope)
	programNode := parser.InitializeParserFor(codeString)
	return evaluator.EvaluateProgramNode(programNode)
}

//real world API
func NewEvaluatorWithGlobalScope(global *treeNodes.Scope) *Evaluator {
	evaluator := new(Evaluator)
	evaluator.programCache = make(map[string]treeNodes.ProgramNodeInterface)
	evaluator.globalScope = global
	return evaluator
}

type Evaluator struct {
	globalScope  *treeNodes.Scope
	programCache map[string]treeNodes.ProgramNodeInterface
}

func (e *Evaluator) SetGlobalScope(scope *treeNodes.Scope) *Evaluator {
	e.globalScope = scope
	return e
}

func (e *Evaluator) RunProgram(programString string) treeNodes.SmalltalkObjectInterface {
	programNode, ok := e.programCache[programString]
	if !ok {
		programNode = parser.InitializeParserFor(programString)
		e.programCache[programString] = programNode
	}
	return e.EvaluateProgramNode(programNode)
}

func (e *Evaluator) EvaluateProgramNode(programNode treeNodes.ProgramNodeInterface) treeNodes.SmalltalkObjectInterface {
	localScope := new(treeNodes.Scope).Initialize()
	localScope.OuterScope = e.globalScope
	return programNode.Eval(localScope)
}