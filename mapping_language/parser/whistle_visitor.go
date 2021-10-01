// Code generated from ./parser/Whistle.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Whistle
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by WhistleParser.
type WhistleVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by WhistleParser#bioperator1.
	VisitBioperator1(ctx *Bioperator1Context) interface{}

	// Visit a parse tree produced by WhistleParser#bioperator2.
	VisitBioperator2(ctx *Bioperator2Context) interface{}

	// Visit a parse tree produced by WhistleParser#bioperator3.
	VisitBioperator3(ctx *Bioperator3Context) interface{}

	// Visit a parse tree produced by WhistleParser#bioperator4.
	VisitBioperator4(ctx *Bioperator4Context) interface{}

	// Visit a parse tree produced by WhistleParser#postunoperator.
	VisitPostunoperator(ctx *PostunoperatorContext) interface{}

	// Visit a parse tree produced by WhistleParser#preunoperator.
	VisitPreunoperator(ctx *PreunoperatorContext) interface{}

	// Visit a parse tree produced by WhistleParser#floatingPoint.
	VisitFloatingPoint(ctx *FloatingPointContext) interface{}

	// Visit a parse tree produced by WhistleParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by WhistleParser#projectorDef.
	VisitProjectorDef(ctx *ProjectorDefContext) interface{}

	// Visit a parse tree produced by WhistleParser#argAlias.
	VisitArgAlias(ctx *ArgAliasContext) interface{}

	// Visit a parse tree produced by WhistleParser#conditionBlock.
	VisitConditionBlock(ctx *ConditionBlockContext) interface{}

	// Visit a parse tree produced by WhistleParser#ifBlock.
	VisitIfBlock(ctx *IfBlockContext) interface{}

	// Visit a parse tree produced by WhistleParser#elseBlock.
	VisitElseBlock(ctx *ElseBlockContext) interface{}

	// Visit a parse tree produced by WhistleParser#inlineCondition.
	VisitInlineCondition(ctx *InlineConditionContext) interface{}

	// Visit a parse tree produced by WhistleParser#inlineFilter.
	VisitInlineFilter(ctx *InlineFilterContext) interface{}

	// Visit a parse tree produced by WhistleParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by WhistleParser#arrayMod.
	VisitArrayMod(ctx *ArrayModContext) interface{}

	// Visit a parse tree produced by WhistleParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by WhistleParser#mapping.
	VisitMapping(ctx *MappingContext) interface{}

	// Visit a parse tree produced by WhistleParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by WhistleParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by WhistleParser#filter.
	VisitFilter(ctx *FilterContext) interface{}

	// Visit a parse tree produced by WhistleParser#ListInitialization.
	VisitListInitialization(ctx *ListInitializationContext) interface{}

	// Visit a parse tree produced by WhistleParser#ExprPreOp.
	VisitExprPreOp(ctx *ExprPreOpContext) interface{}

	// Visit a parse tree produced by WhistleParser#ExprPostOp.
	VisitExprPostOp(ctx *ExprPostOpContext) interface{}

	// Visit a parse tree produced by WhistleParser#ExprAnonBlock.
	VisitExprAnonBlock(ctx *ExprAnonBlockContext) interface{}

	// Visit a parse tree produced by WhistleParser#ExprProjection.
	VisitExprProjection(ctx *ExprProjectionContext) interface{}

	// Visit a parse tree produced by WhistleParser#ExprBiOp.
	VisitExprBiOp(ctx *ExprBiOpContext) interface{}

	// Visit a parse tree produced by WhistleParser#ExprSource.
	VisitExprSource(ctx *ExprSourceContext) interface{}

	// Visit a parse tree produced by WhistleParser#SourceConstNum.
	VisitSourceConstNum(ctx *SourceConstNumContext) interface{}

	// Visit a parse tree produced by WhistleParser#SourceInput.
	VisitSourceInput(ctx *SourceInputContext) interface{}

	// Visit a parse tree produced by WhistleParser#SourceConstStr.
	VisitSourceConstStr(ctx *SourceConstStrContext) interface{}

	// Visit a parse tree produced by WhistleParser#SourceConstBool.
	VisitSourceConstBool(ctx *SourceConstBoolContext) interface{}

	// Visit a parse tree produced by WhistleParser#SourceProjection.
	VisitSourceProjection(ctx *SourceProjectionContext) interface{}

	// Visit a parse tree produced by WhistleParser#TargetVar.
	VisitTargetVar(ctx *TargetVarContext) interface{}

	// Visit a parse tree produced by WhistleParser#TargetRootField.
	VisitTargetRootField(ctx *TargetRootFieldContext) interface{}

	// Visit a parse tree produced by WhistleParser#TargetObj.
	VisitTargetObj(ctx *TargetObjContext) interface{}

	// Visit a parse tree produced by WhistleParser#TargetThis.
	VisitTargetThis(ctx *TargetThisContext) interface{}

	// Visit a parse tree produced by WhistleParser#TargetField.
	VisitTargetField(ctx *TargetFieldContext) interface{}

	// Visit a parse tree produced by WhistleParser#targetPath.
	VisitTargetPath(ctx *TargetPathContext) interface{}

	// Visit a parse tree produced by WhistleParser#targetPathHead.
	VisitTargetPathHead(ctx *TargetPathHeadContext) interface{}

	// Visit a parse tree produced by WhistleParser#targetPathSegment.
	VisitTargetPathSegment(ctx *TargetPathSegmentContext) interface{}

	// Visit a parse tree produced by WhistleParser#sourcePath.
	VisitSourcePath(ctx *SourcePathContext) interface{}

	// Visit a parse tree produced by WhistleParser#sourcePathHead.
	VisitSourcePathHead(ctx *SourcePathHeadContext) interface{}

	// Visit a parse tree produced by WhistleParser#sourcePathSegment.
	VisitSourcePathSegment(ctx *SourcePathSegmentContext) interface{}

	// Visit a parse tree produced by WhistleParser#postProcessInline.
	VisitPostProcessInline(ctx *PostProcessInlineContext) interface{}

	// Visit a parse tree produced by WhistleParser#postProcessName.
	VisitPostProcessName(ctx *PostProcessNameContext) interface{}
}
