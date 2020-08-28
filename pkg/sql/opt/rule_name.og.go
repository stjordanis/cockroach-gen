// Code generated by optgen; DO NOT EDIT.

package opt

const (
	startAutoRule RuleName = iota + NumManualRuleNames

	// ------------------------------------------------------------
	// Normalize Rule Names
	// ------------------------------------------------------------
	EliminateAggDistinct
	NormalizeNestedAnds
	SimplifyTrueAnd
	SimplifyAndTrue
	SimplifyFalseAnd
	SimplifyAndFalse
	SimplifyTrueOr
	SimplifyOrTrue
	SimplifyFalseOr
	SimplifyOrFalse
	SimplifyRange
	FoldNullAndOr
	FoldNotTrue
	FoldNotFalse
	FoldNotNull
	NegateComparison
	EliminateNot
	NegateAnd
	NegateOr
	ExtractRedundantConjunct
	CommuteVarInequality
	CommuteConstInequality
	NormalizeCmpPlusConst
	NormalizeCmpMinusConst
	NormalizeCmpConstMinus
	NormalizeTupleEquality
	FoldNullComparisonLeft
	FoldNullComparisonRight
	FoldIsNull
	FoldNonNullIsNull
	FoldNullTupleIsTupleNull
	FoldNonNullTupleIsTupleNull
	FoldIsNotNull
	FoldNonNullIsNotNull
	FoldNonNullTupleIsTupleNotNull
	FoldNullTupleIsTupleNotNull
	CommuteNullIs
	NormalizeCmpTimeZoneFunction
	NormalizeCmpTimeZoneFunctionTZ
	FoldEqZeroSTDistance
	FoldCmpSTDistanceLeft
	FoldCmpSTDistanceRight
	FoldCmpSTMaxDistanceLeft
	FoldCmpSTMaxDistanceRight
	DecorrelateJoin
	DecorrelateProjectSet
	TryDecorrelateSelect
	TryDecorrelateProject
	TryDecorrelateProjectSelect
	TryDecorrelateProjectInnerJoin
	TryDecorrelateInnerJoin
	TryDecorrelateInnerLeftJoin
	TryDecorrelateGroupBy
	TryDecorrelateScalarGroupBy
	TryDecorrelateSemiJoin
	TryDecorrelateLimitOne
	TryDecorrelateProjectSet
	TryDecorrelateWindow
	TryDecorrelateMax1Row
	HoistSelectExists
	HoistSelectNotExists
	HoistSelectSubquery
	HoistProjectSubquery
	HoistJoinSubquery
	HoistValuesSubquery
	HoistProjectSetSubquery
	NormalizeSelectAnyFilter
	NormalizeJoinAnyFilter
	NormalizeSelectNotAnyFilter
	NormalizeJoinNotAnyFilter
	FoldNullCast
	FoldNullUnary
	FoldNullBinaryLeft
	FoldNullBinaryRight
	FoldNullInNonEmpty
	FoldInEmpty
	FoldNotInEmpty
	FoldArray
	FoldBinary
	FoldUnary
	FoldComparison
	FoldCast
	FoldIndirection
	FoldColumnAccess
	FoldFunction
	FoldEqualsAnyNull
	ConvertGroupByToDistinct
	EliminateGroupByProject
	EliminateJoinUnderGroupByLeft
	EliminateJoinUnderGroupByRight
	EliminateDistinct
	ReduceGroupingCols
	ReduceNotNullGroupingCols
	EliminateAggDistinctForKeys
	EliminateAggFilteredDistinctForKeys
	EliminateDistinctNoColumns
	EliminateEnsureDistinctNoColumns
	EliminateDistinctOnValues
	PushAggDistinctIntoGroupBy
	PushAggFilterIntoScalarGroupBy
	ConvertCountToCountRows
	FoldGroupingOperators
	InlineConstVar
	InlineProjectConstants
	InlineSelectConstants
	InlineJoinConstantsLeft
	InlineJoinConstantsRight
	PushSelectIntoInlinableProject
	InlineProjectInProject
	CommuteRightJoin
	SimplifyJoinFilters
	DetectJoinContradiction
	PushFilterIntoJoinLeftAndRight
	MapFilterIntoJoinLeft
	MapFilterIntoJoinRight
	MapEqualityIntoJoinLeftAndRight
	PushFilterIntoJoinLeft
	PushFilterIntoJoinRight
	SimplifyLeftJoin
	SimplifyRightJoin
	EliminateSemiJoin
	SimplifyZeroCardinalitySemiJoin
	EliminateAntiJoin
	SimplifyZeroCardinalityAntiJoin
	EliminateJoinNoColsLeft
	EliminateJoinNoColsRight
	HoistJoinProjectRight
	HoistJoinProjectLeft
	SimplifyJoinNotNullEquality
	ExtractJoinEqualities
	SortFiltersInJoin
	LeftAssociateJoinsLeft
	LeftAssociateJoinsRight
	RightAssociateJoinsLeft
	RightAssociateJoinsRight
	RemoveJoinNotNullCondition
	EliminateLimit
	EliminateOffset
	PushLimitIntoProject
	PushOffsetIntoProject
	PushLimitIntoOffset
	PushLimitIntoOrdinality
	PushLimitIntoJoinLeft
	PushLimitIntoJoinRight
	FoldLimits
	AssociateLimitJoinsLeft
	AssociateLimitJoinsRight
	EliminateMax1Row
	FoldPlusZero
	FoldZeroPlus
	FoldMinusZero
	FoldMultOne
	FoldOneMult
	FoldDivOne
	InvertMinus
	EliminateUnaryMinus
	SimplifyLimitOrdering
	SimplifyOffsetOrdering
	SimplifyGroupByOrdering
	SimplifyOrdinalityOrdering
	SimplifyExplainOrdering
	EliminateJoinUnderProjectLeft
	EliminateJoinUnderProjectRight
	EliminateProject
	MergeProjects
	MergeProjectWithValues
	PushColumnRemappingIntoValues
	FoldTupleAccessIntoValues
	FoldJSONAccessIntoValues
	ConvertZipArraysToValues
	PruneProjectCols
	PruneScanCols
	PruneSelectCols
	PruneLimitCols
	PruneOffsetCols
	PruneJoinLeftCols
	PruneJoinRightCols
	PruneSemiAntiJoinRightCols
	PruneAggCols
	PruneGroupByCols
	PruneValuesCols
	PruneOrdinalityCols
	PruneExplainCols
	PruneProjectSetCols
	PruneWindowOutputCols
	PruneWindowInputCols
	PruneMutationFetchCols
	PruneMutationInputCols
	PruneMutationReturnCols
	PruneWithScanCols
	PruneWithCols
	PruneUnionAllCols
	RejectNullsLeftJoin
	RejectNullsRightJoin
	RejectNullsGroupBy
	RejectNullsUnderJoinLeft
	RejectNullsUnderJoinRight
	RejectNullsProject
	CommuteVar
	CommuteConst
	EliminateCoalesce
	SimplifyCoalesce
	EliminateCast
	NormalizeInConst
	FoldInNull
	UnifyComparisonTypes
	EliminateExistsZeroRows
	EliminateExistsProject
	EliminateExistsGroupBy
	IntroduceExistsLimit
	EliminateExistsLimit
	NormalizeJSONFieldAccess
	NormalizeJSONContains
	SimplifyCaseWhenConstValue
	InlineAnyValuesSingleCol
	InlineAnyValuesMultiCol
	SimplifyEqualsAnyTuple
	SimplifyAnyScalarArray
	FoldCollate
	NormalizeArrayFlattenToAgg
	SimplifySameVarEqualities
	SimplifySameVarInequalities
	SimplifySelectFilters
	ConsolidateSelectFilters
	DetectSelectContradiction
	EliminateSelect
	MergeSelects
	PushSelectIntoProject
	MergeSelectInnerJoin
	PushSelectCondLeftIntoJoinLeftAndRight
	PushSelectIntoJoinLeft
	PushSelectIntoGroupBy
	RemoveNotNullCondition
	PushSelectIntoProjectSet
	PushFilterIntoSetOp
	EliminateUnionAllLeft
	EliminateUnionAllRight
	EliminateWindow
	ReduceWindowPartitionCols
	SimplifyWindowOrdering
	PushSelectIntoWindow
	PushLimitIntoWindow
	InlineWith

	// startExploreRule tracks the number of normalization rules;
	// all rules greater than this value are exploration rules.
	startExploreRule

	// ------------------------------------------------------------
	// Explore Rule Names
	// ------------------------------------------------------------
	ReplaceScalarMinMaxWithLimit
	ReplaceMinWithLimit
	ReplaceMaxWithLimit
	GenerateStreamingGroupBy
	ReorderJoins
	CommuteLeftJoin
	CommuteSemiJoin
	ConvertSemiToInnerJoin
	GenerateMergeJoins
	GenerateLookupJoins
	GenerateInvertedJoins
	GenerateInvertedJoinsFromSelect
	GenerateZigzagJoins
	GenerateInvertedIndexZigzagJoins
	GenerateLookupJoinsWithFilter
	PushJoinIntoIndexJoin
	GenerateLimitedScans
	PushLimitIntoFilteredScan
	PushLimitIntoIndexJoin
	SplitScanIntoUnionScans
	GenerateIndexScans
	GeneratePartialIndexScans
	GenerateConstrainedScans
	GenerateInvertedIndexScans
	SplitDisjunction
	SplitDisjunctionAddKey

	// NumRuleNames tracks the total count of rule names.
	NumRuleNames
)
