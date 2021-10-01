// Code generated from ./parser/Whistle.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Whistle
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWhistleListener is a complete listener for a parse tree produced by WhistleParser.
type BaseWhistleListener struct{}

var _ WhistleListener = &BaseWhistleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWhistleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWhistleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWhistleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWhistleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterBioperator1 is called when production bioperator1 is entered.
func (s *BaseWhistleListener) EnterBioperator1(ctx *Bioperator1Context) {}

// ExitBioperator1 is called when production bioperator1 is exited.
func (s *BaseWhistleListener) ExitBioperator1(ctx *Bioperator1Context) {}

// EnterBioperator2 is called when production bioperator2 is entered.
func (s *BaseWhistleListener) EnterBioperator2(ctx *Bioperator2Context) {}

// ExitBioperator2 is called when production bioperator2 is exited.
func (s *BaseWhistleListener) ExitBioperator2(ctx *Bioperator2Context) {}

// EnterBioperator3 is called when production bioperator3 is entered.
func (s *BaseWhistleListener) EnterBioperator3(ctx *Bioperator3Context) {}

// ExitBioperator3 is called when production bioperator3 is exited.
func (s *BaseWhistleListener) ExitBioperator3(ctx *Bioperator3Context) {}

// EnterBioperator4 is called when production bioperator4 is entered.
func (s *BaseWhistleListener) EnterBioperator4(ctx *Bioperator4Context) {}

// ExitBioperator4 is called when production bioperator4 is exited.
func (s *BaseWhistleListener) ExitBioperator4(ctx *Bioperator4Context) {}

// EnterPostunoperator is called when production postunoperator is entered.
func (s *BaseWhistleListener) EnterPostunoperator(ctx *PostunoperatorContext) {}

// ExitPostunoperator is called when production postunoperator is exited.
func (s *BaseWhistleListener) ExitPostunoperator(ctx *PostunoperatorContext) {}

// EnterPreunoperator is called when production preunoperator is entered.
func (s *BaseWhistleListener) EnterPreunoperator(ctx *PreunoperatorContext) {}

// ExitPreunoperator is called when production preunoperator is exited.
func (s *BaseWhistleListener) ExitPreunoperator(ctx *PreunoperatorContext) {}

// EnterFloatingPoint is called when production floatingPoint is entered.
func (s *BaseWhistleListener) EnterFloatingPoint(ctx *FloatingPointContext) {}

// ExitFloatingPoint is called when production floatingPoint is exited.
func (s *BaseWhistleListener) ExitFloatingPoint(ctx *FloatingPointContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseWhistleListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseWhistleListener) ExitRoot(ctx *RootContext) {}

// EnterProjectorDef is called when production projectorDef is entered.
func (s *BaseWhistleListener) EnterProjectorDef(ctx *ProjectorDefContext) {}

// ExitProjectorDef is called when production projectorDef is exited.
func (s *BaseWhistleListener) ExitProjectorDef(ctx *ProjectorDefContext) {}

// EnterArgAlias is called when production argAlias is entered.
func (s *BaseWhistleListener) EnterArgAlias(ctx *ArgAliasContext) {}

// ExitArgAlias is called when production argAlias is exited.
func (s *BaseWhistleListener) ExitArgAlias(ctx *ArgAliasContext) {}

// EnterConditionBlock is called when production conditionBlock is entered.
func (s *BaseWhistleListener) EnterConditionBlock(ctx *ConditionBlockContext) {}

// ExitConditionBlock is called when production conditionBlock is exited.
func (s *BaseWhistleListener) ExitConditionBlock(ctx *ConditionBlockContext) {}

// EnterIfBlock is called when production ifBlock is entered.
func (s *BaseWhistleListener) EnterIfBlock(ctx *IfBlockContext) {}

// ExitIfBlock is called when production ifBlock is exited.
func (s *BaseWhistleListener) ExitIfBlock(ctx *IfBlockContext) {}

// EnterElseBlock is called when production elseBlock is entered.
func (s *BaseWhistleListener) EnterElseBlock(ctx *ElseBlockContext) {}

// ExitElseBlock is called when production elseBlock is exited.
func (s *BaseWhistleListener) ExitElseBlock(ctx *ElseBlockContext) {}

// EnterInlineCondition is called when production inlineCondition is entered.
func (s *BaseWhistleListener) EnterInlineCondition(ctx *InlineConditionContext) {}

// ExitInlineCondition is called when production inlineCondition is exited.
func (s *BaseWhistleListener) ExitInlineCondition(ctx *InlineConditionContext) {}

// EnterInlineFilter is called when production inlineFilter is entered.
func (s *BaseWhistleListener) EnterInlineFilter(ctx *InlineFilterContext) {}

// ExitInlineFilter is called when production inlineFilter is exited.
func (s *BaseWhistleListener) ExitInlineFilter(ctx *InlineFilterContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseWhistleListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseWhistleListener) ExitIndex(ctx *IndexContext) {}

// EnterArrayMod is called when production arrayMod is entered.
func (s *BaseWhistleListener) EnterArrayMod(ctx *ArrayModContext) {}

// ExitArrayMod is called when production arrayMod is exited.
func (s *BaseWhistleListener) ExitArrayMod(ctx *ArrayModContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseWhistleListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseWhistleListener) ExitBlock(ctx *BlockContext) {}

// EnterMapping is called when production mapping is entered.
func (s *BaseWhistleListener) EnterMapping(ctx *MappingContext) {}

// ExitMapping is called when production mapping is exited.
func (s *BaseWhistleListener) ExitMapping(ctx *MappingContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseWhistleListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseWhistleListener) ExitComment(ctx *CommentContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseWhistleListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseWhistleListener) ExitCondition(ctx *ConditionContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseWhistleListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseWhistleListener) ExitFilter(ctx *FilterContext) {}

// EnterListInitialization is called when production ListInitialization is entered.
func (s *BaseWhistleListener) EnterListInitialization(ctx *ListInitializationContext) {}

// ExitListInitialization is called when production ListInitialization is exited.
func (s *BaseWhistleListener) ExitListInitialization(ctx *ListInitializationContext) {}

// EnterExprPreOp is called when production ExprPreOp is entered.
func (s *BaseWhistleListener) EnterExprPreOp(ctx *ExprPreOpContext) {}

// ExitExprPreOp is called when production ExprPreOp is exited.
func (s *BaseWhistleListener) ExitExprPreOp(ctx *ExprPreOpContext) {}

// EnterExprPostOp is called when production ExprPostOp is entered.
func (s *BaseWhistleListener) EnterExprPostOp(ctx *ExprPostOpContext) {}

// ExitExprPostOp is called when production ExprPostOp is exited.
func (s *BaseWhistleListener) ExitExprPostOp(ctx *ExprPostOpContext) {}

// EnterExprAnonBlock is called when production ExprAnonBlock is entered.
func (s *BaseWhistleListener) EnterExprAnonBlock(ctx *ExprAnonBlockContext) {}

// ExitExprAnonBlock is called when production ExprAnonBlock is exited.
func (s *BaseWhistleListener) ExitExprAnonBlock(ctx *ExprAnonBlockContext) {}

// EnterExprProjection is called when production ExprProjection is entered.
func (s *BaseWhistleListener) EnterExprProjection(ctx *ExprProjectionContext) {}

// ExitExprProjection is called when production ExprProjection is exited.
func (s *BaseWhistleListener) ExitExprProjection(ctx *ExprProjectionContext) {}

// EnterExprBiOp is called when production ExprBiOp is entered.
func (s *BaseWhistleListener) EnterExprBiOp(ctx *ExprBiOpContext) {}

// ExitExprBiOp is called when production ExprBiOp is exited.
func (s *BaseWhistleListener) ExitExprBiOp(ctx *ExprBiOpContext) {}

// EnterExprSource is called when production ExprSource is entered.
func (s *BaseWhistleListener) EnterExprSource(ctx *ExprSourceContext) {}

// ExitExprSource is called when production ExprSource is exited.
func (s *BaseWhistleListener) ExitExprSource(ctx *ExprSourceContext) {}

// EnterSourceConstNum is called when production SourceConstNum is entered.
func (s *BaseWhistleListener) EnterSourceConstNum(ctx *SourceConstNumContext) {}

// ExitSourceConstNum is called when production SourceConstNum is exited.
func (s *BaseWhistleListener) ExitSourceConstNum(ctx *SourceConstNumContext) {}

// EnterSourceInput is called when production SourceInput is entered.
func (s *BaseWhistleListener) EnterSourceInput(ctx *SourceInputContext) {}

// ExitSourceInput is called when production SourceInput is exited.
func (s *BaseWhistleListener) ExitSourceInput(ctx *SourceInputContext) {}

// EnterSourceConstStr is called when production SourceConstStr is entered.
func (s *BaseWhistleListener) EnterSourceConstStr(ctx *SourceConstStrContext) {}

// ExitSourceConstStr is called when production SourceConstStr is exited.
func (s *BaseWhistleListener) ExitSourceConstStr(ctx *SourceConstStrContext) {}

// EnterSourceConstBool is called when production SourceConstBool is entered.
func (s *BaseWhistleListener) EnterSourceConstBool(ctx *SourceConstBoolContext) {}

// ExitSourceConstBool is called when production SourceConstBool is exited.
func (s *BaseWhistleListener) ExitSourceConstBool(ctx *SourceConstBoolContext) {}

// EnterSourceProjection is called when production SourceProjection is entered.
func (s *BaseWhistleListener) EnterSourceProjection(ctx *SourceProjectionContext) {}

// ExitSourceProjection is called when production SourceProjection is exited.
func (s *BaseWhistleListener) ExitSourceProjection(ctx *SourceProjectionContext) {}

// EnterTargetVar is called when production TargetVar is entered.
func (s *BaseWhistleListener) EnterTargetVar(ctx *TargetVarContext) {}

// ExitTargetVar is called when production TargetVar is exited.
func (s *BaseWhistleListener) ExitTargetVar(ctx *TargetVarContext) {}

// EnterTargetRootField is called when production TargetRootField is entered.
func (s *BaseWhistleListener) EnterTargetRootField(ctx *TargetRootFieldContext) {}

// ExitTargetRootField is called when production TargetRootField is exited.
func (s *BaseWhistleListener) ExitTargetRootField(ctx *TargetRootFieldContext) {}

// EnterTargetObj is called when production TargetObj is entered.
func (s *BaseWhistleListener) EnterTargetObj(ctx *TargetObjContext) {}

// ExitTargetObj is called when production TargetObj is exited.
func (s *BaseWhistleListener) ExitTargetObj(ctx *TargetObjContext) {}

// EnterTargetThis is called when production TargetThis is entered.
func (s *BaseWhistleListener) EnterTargetThis(ctx *TargetThisContext) {}

// ExitTargetThis is called when production TargetThis is exited.
func (s *BaseWhistleListener) ExitTargetThis(ctx *TargetThisContext) {}

// EnterTargetField is called when production TargetField is entered.
func (s *BaseWhistleListener) EnterTargetField(ctx *TargetFieldContext) {}

// ExitTargetField is called when production TargetField is exited.
func (s *BaseWhistleListener) ExitTargetField(ctx *TargetFieldContext) {}

// EnterTargetPath is called when production targetPath is entered.
func (s *BaseWhistleListener) EnterTargetPath(ctx *TargetPathContext) {}

// ExitTargetPath is called when production targetPath is exited.
func (s *BaseWhistleListener) ExitTargetPath(ctx *TargetPathContext) {}

// EnterTargetPathHead is called when production targetPathHead is entered.
func (s *BaseWhistleListener) EnterTargetPathHead(ctx *TargetPathHeadContext) {}

// ExitTargetPathHead is called when production targetPathHead is exited.
func (s *BaseWhistleListener) ExitTargetPathHead(ctx *TargetPathHeadContext) {}

// EnterTargetPathSegment is called when production targetPathSegment is entered.
func (s *BaseWhistleListener) EnterTargetPathSegment(ctx *TargetPathSegmentContext) {}

// ExitTargetPathSegment is called when production targetPathSegment is exited.
func (s *BaseWhistleListener) ExitTargetPathSegment(ctx *TargetPathSegmentContext) {}

// EnterSourcePath is called when production sourcePath is entered.
func (s *BaseWhistleListener) EnterSourcePath(ctx *SourcePathContext) {}

// ExitSourcePath is called when production sourcePath is exited.
func (s *BaseWhistleListener) ExitSourcePath(ctx *SourcePathContext) {}

// EnterSourcePathHead is called when production sourcePathHead is entered.
func (s *BaseWhistleListener) EnterSourcePathHead(ctx *SourcePathHeadContext) {}

// ExitSourcePathHead is called when production sourcePathHead is exited.
func (s *BaseWhistleListener) ExitSourcePathHead(ctx *SourcePathHeadContext) {}

// EnterSourcePathSegment is called when production sourcePathSegment is entered.
func (s *BaseWhistleListener) EnterSourcePathSegment(ctx *SourcePathSegmentContext) {}

// ExitSourcePathSegment is called when production sourcePathSegment is exited.
func (s *BaseWhistleListener) ExitSourcePathSegment(ctx *SourcePathSegmentContext) {}

// EnterPostProcessInline is called when production postProcessInline is entered.
func (s *BaseWhistleListener) EnterPostProcessInline(ctx *PostProcessInlineContext) {}

// ExitPostProcessInline is called when production postProcessInline is exited.
func (s *BaseWhistleListener) ExitPostProcessInline(ctx *PostProcessInlineContext) {}

// EnterPostProcessName is called when production postProcessName is entered.
func (s *BaseWhistleListener) EnterPostProcessName(ctx *PostProcessNameContext) {}

// ExitPostProcessName is called when production postProcessName is exited.
func (s *BaseWhistleListener) ExitPostProcessName(ctx *PostProcessNameContext) {}
