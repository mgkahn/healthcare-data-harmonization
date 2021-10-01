// Code generated from ./parser/Whistle.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Whistle
import "github.com/antlr/antlr4/runtime/Go/antlr"

// WhistleListener is a complete listener for a parse tree produced by WhistleParser.
type WhistleListener interface {
	antlr.ParseTreeListener

	// EnterBioperator1 is called when entering the bioperator1 production.
	EnterBioperator1(c *Bioperator1Context)

	// EnterBioperator2 is called when entering the bioperator2 production.
	EnterBioperator2(c *Bioperator2Context)

	// EnterBioperator3 is called when entering the bioperator3 production.
	EnterBioperator3(c *Bioperator3Context)

	// EnterBioperator4 is called when entering the bioperator4 production.
	EnterBioperator4(c *Bioperator4Context)

	// EnterPostunoperator is called when entering the postunoperator production.
	EnterPostunoperator(c *PostunoperatorContext)

	// EnterPreunoperator is called when entering the preunoperator production.
	EnterPreunoperator(c *PreunoperatorContext)

	// EnterFloatingPoint is called when entering the floatingPoint production.
	EnterFloatingPoint(c *FloatingPointContext)

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterProjectorDef is called when entering the projectorDef production.
	EnterProjectorDef(c *ProjectorDefContext)

	// EnterArgAlias is called when entering the argAlias production.
	EnterArgAlias(c *ArgAliasContext)

	// EnterConditionBlock is called when entering the conditionBlock production.
	EnterConditionBlock(c *ConditionBlockContext)

	// EnterIfBlock is called when entering the ifBlock production.
	EnterIfBlock(c *IfBlockContext)

	// EnterElseBlock is called when entering the elseBlock production.
	EnterElseBlock(c *ElseBlockContext)

	// EnterInlineCondition is called when entering the inlineCondition production.
	EnterInlineCondition(c *InlineConditionContext)

	// EnterInlineFilter is called when entering the inlineFilter production.
	EnterInlineFilter(c *InlineFilterContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterArrayMod is called when entering the arrayMod production.
	EnterArrayMod(c *ArrayModContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterMapping is called when entering the mapping production.
	EnterMapping(c *MappingContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterListInitialization is called when entering the ListInitialization production.
	EnterListInitialization(c *ListInitializationContext)

	// EnterExprPreOp is called when entering the ExprPreOp production.
	EnterExprPreOp(c *ExprPreOpContext)

	// EnterExprPostOp is called when entering the ExprPostOp production.
	EnterExprPostOp(c *ExprPostOpContext)

	// EnterExprAnonBlock is called when entering the ExprAnonBlock production.
	EnterExprAnonBlock(c *ExprAnonBlockContext)

	// EnterExprProjection is called when entering the ExprProjection production.
	EnterExprProjection(c *ExprProjectionContext)

	// EnterExprBiOp is called when entering the ExprBiOp production.
	EnterExprBiOp(c *ExprBiOpContext)

	// EnterExprSource is called when entering the ExprSource production.
	EnterExprSource(c *ExprSourceContext)

	// EnterSourceConstNum is called when entering the SourceConstNum production.
	EnterSourceConstNum(c *SourceConstNumContext)

	// EnterSourceInput is called when entering the SourceInput production.
	EnterSourceInput(c *SourceInputContext)

	// EnterSourceConstStr is called when entering the SourceConstStr production.
	EnterSourceConstStr(c *SourceConstStrContext)

	// EnterSourceConstBool is called when entering the SourceConstBool production.
	EnterSourceConstBool(c *SourceConstBoolContext)

	// EnterSourceProjection is called when entering the SourceProjection production.
	EnterSourceProjection(c *SourceProjectionContext)

	// EnterTargetVar is called when entering the TargetVar production.
	EnterTargetVar(c *TargetVarContext)

	// EnterTargetRootField is called when entering the TargetRootField production.
	EnterTargetRootField(c *TargetRootFieldContext)

	// EnterTargetObj is called when entering the TargetObj production.
	EnterTargetObj(c *TargetObjContext)

	// EnterTargetThis is called when entering the TargetThis production.
	EnterTargetThis(c *TargetThisContext)

	// EnterTargetField is called when entering the TargetField production.
	EnterTargetField(c *TargetFieldContext)

	// EnterTargetPath is called when entering the targetPath production.
	EnterTargetPath(c *TargetPathContext)

	// EnterTargetPathHead is called when entering the targetPathHead production.
	EnterTargetPathHead(c *TargetPathHeadContext)

	// EnterTargetPathSegment is called when entering the targetPathSegment production.
	EnterTargetPathSegment(c *TargetPathSegmentContext)

	// EnterSourcePath is called when entering the sourcePath production.
	EnterSourcePath(c *SourcePathContext)

	// EnterSourcePathHead is called when entering the sourcePathHead production.
	EnterSourcePathHead(c *SourcePathHeadContext)

	// EnterSourcePathSegment is called when entering the sourcePathSegment production.
	EnterSourcePathSegment(c *SourcePathSegmentContext)

	// EnterPostProcessInline is called when entering the postProcessInline production.
	EnterPostProcessInline(c *PostProcessInlineContext)

	// EnterPostProcessName is called when entering the postProcessName production.
	EnterPostProcessName(c *PostProcessNameContext)

	// ExitBioperator1 is called when exiting the bioperator1 production.
	ExitBioperator1(c *Bioperator1Context)

	// ExitBioperator2 is called when exiting the bioperator2 production.
	ExitBioperator2(c *Bioperator2Context)

	// ExitBioperator3 is called when exiting the bioperator3 production.
	ExitBioperator3(c *Bioperator3Context)

	// ExitBioperator4 is called when exiting the bioperator4 production.
	ExitBioperator4(c *Bioperator4Context)

	// ExitPostunoperator is called when exiting the postunoperator production.
	ExitPostunoperator(c *PostunoperatorContext)

	// ExitPreunoperator is called when exiting the preunoperator production.
	ExitPreunoperator(c *PreunoperatorContext)

	// ExitFloatingPoint is called when exiting the floatingPoint production.
	ExitFloatingPoint(c *FloatingPointContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitProjectorDef is called when exiting the projectorDef production.
	ExitProjectorDef(c *ProjectorDefContext)

	// ExitArgAlias is called when exiting the argAlias production.
	ExitArgAlias(c *ArgAliasContext)

	// ExitConditionBlock is called when exiting the conditionBlock production.
	ExitConditionBlock(c *ConditionBlockContext)

	// ExitIfBlock is called when exiting the ifBlock production.
	ExitIfBlock(c *IfBlockContext)

	// ExitElseBlock is called when exiting the elseBlock production.
	ExitElseBlock(c *ElseBlockContext)

	// ExitInlineCondition is called when exiting the inlineCondition production.
	ExitInlineCondition(c *InlineConditionContext)

	// ExitInlineFilter is called when exiting the inlineFilter production.
	ExitInlineFilter(c *InlineFilterContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitArrayMod is called when exiting the arrayMod production.
	ExitArrayMod(c *ArrayModContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitMapping is called when exiting the mapping production.
	ExitMapping(c *MappingContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitListInitialization is called when exiting the ListInitialization production.
	ExitListInitialization(c *ListInitializationContext)

	// ExitExprPreOp is called when exiting the ExprPreOp production.
	ExitExprPreOp(c *ExprPreOpContext)

	// ExitExprPostOp is called when exiting the ExprPostOp production.
	ExitExprPostOp(c *ExprPostOpContext)

	// ExitExprAnonBlock is called when exiting the ExprAnonBlock production.
	ExitExprAnonBlock(c *ExprAnonBlockContext)

	// ExitExprProjection is called when exiting the ExprProjection production.
	ExitExprProjection(c *ExprProjectionContext)

	// ExitExprBiOp is called when exiting the ExprBiOp production.
	ExitExprBiOp(c *ExprBiOpContext)

	// ExitExprSource is called when exiting the ExprSource production.
	ExitExprSource(c *ExprSourceContext)

	// ExitSourceConstNum is called when exiting the SourceConstNum production.
	ExitSourceConstNum(c *SourceConstNumContext)

	// ExitSourceInput is called when exiting the SourceInput production.
	ExitSourceInput(c *SourceInputContext)

	// ExitSourceConstStr is called when exiting the SourceConstStr production.
	ExitSourceConstStr(c *SourceConstStrContext)

	// ExitSourceConstBool is called when exiting the SourceConstBool production.
	ExitSourceConstBool(c *SourceConstBoolContext)

	// ExitSourceProjection is called when exiting the SourceProjection production.
	ExitSourceProjection(c *SourceProjectionContext)

	// ExitTargetVar is called when exiting the TargetVar production.
	ExitTargetVar(c *TargetVarContext)

	// ExitTargetRootField is called when exiting the TargetRootField production.
	ExitTargetRootField(c *TargetRootFieldContext)

	// ExitTargetObj is called when exiting the TargetObj production.
	ExitTargetObj(c *TargetObjContext)

	// ExitTargetThis is called when exiting the TargetThis production.
	ExitTargetThis(c *TargetThisContext)

	// ExitTargetField is called when exiting the TargetField production.
	ExitTargetField(c *TargetFieldContext)

	// ExitTargetPath is called when exiting the targetPath production.
	ExitTargetPath(c *TargetPathContext)

	// ExitTargetPathHead is called when exiting the targetPathHead production.
	ExitTargetPathHead(c *TargetPathHeadContext)

	// ExitTargetPathSegment is called when exiting the targetPathSegment production.
	ExitTargetPathSegment(c *TargetPathSegmentContext)

	// ExitSourcePath is called when exiting the sourcePath production.
	ExitSourcePath(c *SourcePathContext)

	// ExitSourcePathHead is called when exiting the sourcePathHead production.
	ExitSourcePathHead(c *SourcePathHeadContext)

	// ExitSourcePathSegment is called when exiting the sourcePathSegment production.
	ExitSourcePathSegment(c *SourcePathSegmentContext)

	// ExitPostProcessInline is called when exiting the postProcessInline production.
	ExitPostProcessInline(c *PostProcessInlineContext)

	// ExitPostProcessName is called when exiting the postProcessName production.
	ExitPostProcessName(c *PostProcessNameContext)
}
