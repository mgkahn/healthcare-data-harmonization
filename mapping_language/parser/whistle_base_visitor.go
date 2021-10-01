// Code generated from ./parser/Whistle.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Whistle
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseWhistleVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseWhistleVisitor) VisitBioperator1(ctx *Bioperator1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitBioperator2(ctx *Bioperator2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitBioperator3(ctx *Bioperator3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitBioperator4(ctx *Bioperator4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitPostunoperator(ctx *PostunoperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitPreunoperator(ctx *PreunoperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitFloatingPoint(ctx *FloatingPointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitProjectorDef(ctx *ProjectorDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitArgAlias(ctx *ArgAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitConditionBlock(ctx *ConditionBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitIfBlock(ctx *IfBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitElseBlock(ctx *ElseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitInlineCondition(ctx *InlineConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitInlineFilter(ctx *InlineFilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitArrayMod(ctx *ArrayModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitMapping(ctx *MappingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitFilter(ctx *FilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitListInitialization(ctx *ListInitializationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitExprPreOp(ctx *ExprPreOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitExprPostOp(ctx *ExprPostOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitExprAnonBlock(ctx *ExprAnonBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitExprProjection(ctx *ExprProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitExprBiOp(ctx *ExprBiOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitExprSource(ctx *ExprSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitSourceConstNum(ctx *SourceConstNumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitSourceInput(ctx *SourceInputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitSourceConstStr(ctx *SourceConstStrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitSourceConstBool(ctx *SourceConstBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitSourceProjection(ctx *SourceProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitTargetVar(ctx *TargetVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitTargetRootField(ctx *TargetRootFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitTargetObj(ctx *TargetObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitTargetThis(ctx *TargetThisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitTargetField(ctx *TargetFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitTargetPath(ctx *TargetPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitTargetPathHead(ctx *TargetPathHeadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitTargetPathSegment(ctx *TargetPathSegmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitSourcePath(ctx *SourcePathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitSourcePathHead(ctx *SourcePathHeadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitSourcePathSegment(ctx *SourcePathSegmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitPostProcessInline(ctx *PostProcessInlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWhistleVisitor) VisitPostProcessName(ctx *PostProcessNameContext) interface{} {
	return v.VisitChildren(ctx)
}
