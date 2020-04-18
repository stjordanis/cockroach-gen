// Code generated by "stringer -output=pkg/sql/opt/rule_name_string.go -type=RuleName pkg/sql/opt/rule_name.go pkg/sql/opt/rule_name.og.go"; DO NOT EDIT.

package opt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidRuleName-0]
	_ = x[SimplifyRootOrdering-1]
	_ = x[PruneRootCols-2]
	_ = x[SimplifyZeroCardinalityGroup-3]
	_ = x[NumManualRuleNames-4]
	_ = x[startAutoRule-4]
	_ = x[EliminateAggDistinct-5]
	_ = x[NormalizeNestedAnds-6]
	_ = x[SimplifyTrueAnd-7]
	_ = x[SimplifyAndTrue-8]
	_ = x[SimplifyFalseAnd-9]
	_ = x[SimplifyAndFalse-10]
	_ = x[SimplifyTrueOr-11]
	_ = x[SimplifyOrTrue-12]
	_ = x[SimplifyFalseOr-13]
	_ = x[SimplifyOrFalse-14]
	_ = x[SimplifyRange-15]
	_ = x[FoldNullAndOr-16]
	_ = x[FoldNotTrue-17]
	_ = x[FoldNotFalse-18]
	_ = x[FoldNotNull-19]
	_ = x[NegateComparison-20]
	_ = x[EliminateNot-21]
	_ = x[NegateAnd-22]
	_ = x[NegateOr-23]
	_ = x[ExtractRedundantConjunct-24]
	_ = x[CommuteVarInequality-25]
	_ = x[CommuteConstInequality-26]
	_ = x[NormalizeCmpPlusConst-27]
	_ = x[NormalizeCmpMinusConst-28]
	_ = x[NormalizeCmpConstMinus-29]
	_ = x[NormalizeTupleEquality-30]
	_ = x[FoldNullComparisonLeft-31]
	_ = x[FoldNullComparisonRight-32]
	_ = x[FoldIsNull-33]
	_ = x[FoldNonNullIsNull-34]
	_ = x[FoldIsNotNull-35]
	_ = x[FoldNonNullIsNotNull-36]
	_ = x[CommuteNullIs-37]
	_ = x[DecorrelateJoin-38]
	_ = x[DecorrelateProjectSet-39]
	_ = x[TryDecorrelateSelect-40]
	_ = x[TryDecorrelateProject-41]
	_ = x[TryDecorrelateProjectSelect-42]
	_ = x[TryDecorrelateProjectInnerJoin-43]
	_ = x[TryDecorrelateInnerJoin-44]
	_ = x[TryDecorrelateInnerLeftJoin-45]
	_ = x[TryDecorrelateGroupBy-46]
	_ = x[TryDecorrelateScalarGroupBy-47]
	_ = x[TryDecorrelateSemiJoin-48]
	_ = x[TryDecorrelateLimitOne-49]
	_ = x[TryDecorrelateProjectSet-50]
	_ = x[TryDecorrelateWindow-51]
	_ = x[HoistSelectExists-52]
	_ = x[HoistSelectNotExists-53]
	_ = x[HoistSelectSubquery-54]
	_ = x[HoistProjectSubquery-55]
	_ = x[HoistJoinSubquery-56]
	_ = x[HoistValuesSubquery-57]
	_ = x[HoistProjectSetSubquery-58]
	_ = x[NormalizeSelectAnyFilter-59]
	_ = x[NormalizeJoinAnyFilter-60]
	_ = x[NormalizeSelectNotAnyFilter-61]
	_ = x[NormalizeJoinNotAnyFilter-62]
	_ = x[FoldNullCast-63]
	_ = x[FoldNullUnary-64]
	_ = x[FoldNullBinaryLeft-65]
	_ = x[FoldNullBinaryRight-66]
	_ = x[FoldNullInNonEmpty-67]
	_ = x[FoldInEmpty-68]
	_ = x[FoldNotInEmpty-69]
	_ = x[FoldArray-70]
	_ = x[FoldBinary-71]
	_ = x[FoldUnary-72]
	_ = x[FoldComparison-73]
	_ = x[FoldCast-74]
	_ = x[FoldIndirection-75]
	_ = x[FoldColumnAccess-76]
	_ = x[FoldFunction-77]
	_ = x[FoldEqualsAnyNull-78]
	_ = x[ConvertGroupByToDistinct-79]
	_ = x[EliminateDistinct-80]
	_ = x[EliminateGroupByProject-81]
	_ = x[ReduceGroupingCols-82]
	_ = x[ReduceNotNullGroupingCols-83]
	_ = x[EliminateAggDistinctForKeys-84]
	_ = x[EliminateAggFilteredDistinctForKeys-85]
	_ = x[EliminateDistinctNoColumns-86]
	_ = x[EliminateErrorDistinctNoColumns-87]
	_ = x[EliminateDistinctOnValues-88]
	_ = x[PushAggDistinctIntoScalarGroupBy-89]
	_ = x[PushAggFilterIntoScalarGroupBy-90]
	_ = x[ConvertCountToCountRows-91]
	_ = x[InlineProjectConstants-92]
	_ = x[InlineSelectConstants-93]
	_ = x[InlineJoinConstantsLeft-94]
	_ = x[InlineJoinConstantsRight-95]
	_ = x[PushSelectIntoInlinableProject-96]
	_ = x[InlineProjectInProject-97]
	_ = x[CommuteRightJoin-98]
	_ = x[SimplifyJoinFilters-99]
	_ = x[DetectJoinContradiction-100]
	_ = x[PushFilterIntoJoinLeftAndRight-101]
	_ = x[MapFilterIntoJoinLeft-102]
	_ = x[MapFilterIntoJoinRight-103]
	_ = x[MapEqualityIntoJoinLeftAndRight-104]
	_ = x[PushFilterIntoJoinLeft-105]
	_ = x[PushFilterIntoJoinRight-106]
	_ = x[SimplifyLeftJoinWithoutFilters-107]
	_ = x[SimplifyRightJoinWithoutFilters-108]
	_ = x[SimplifyLeftJoinWithFilters-109]
	_ = x[SimplifyRightJoinWithFilters-110]
	_ = x[EliminateSemiJoin-111]
	_ = x[SimplifyZeroCardinalitySemiJoin-112]
	_ = x[EliminateAntiJoin-113]
	_ = x[SimplifyZeroCardinalityAntiJoin-114]
	_ = x[EliminateJoinNoColsLeft-115]
	_ = x[EliminateJoinNoColsRight-116]
	_ = x[HoistJoinProjectRight-117]
	_ = x[HoistJoinProjectLeft-118]
	_ = x[SimplifyJoinNotNullEquality-119]
	_ = x[ExtractJoinEqualities-120]
	_ = x[SortFiltersInJoin-121]
	_ = x[EliminateLimit-122]
	_ = x[EliminateOffset-123]
	_ = x[PushLimitIntoProject-124]
	_ = x[PushOffsetIntoProject-125]
	_ = x[PushLimitIntoOffset-126]
	_ = x[PushLimitIntoOrdinality-127]
	_ = x[PushLimitIntoLeftJoin-128]
	_ = x[EliminateMax1Row-129]
	_ = x[FoldPlusZero-130]
	_ = x[FoldZeroPlus-131]
	_ = x[FoldMinusZero-132]
	_ = x[FoldMultOne-133]
	_ = x[FoldOneMult-134]
	_ = x[FoldDivOne-135]
	_ = x[InvertMinus-136]
	_ = x[EliminateUnaryMinus-137]
	_ = x[SimplifyLimitOrdering-138]
	_ = x[SimplifyOffsetOrdering-139]
	_ = x[SimplifyGroupByOrdering-140]
	_ = x[SimplifyOrdinalityOrdering-141]
	_ = x[SimplifyExplainOrdering-142]
	_ = x[EliminateProject-143]
	_ = x[MergeProjects-144]
	_ = x[MergeProjectWithValues-145]
	_ = x[PruneProjectCols-146]
	_ = x[PruneScanCols-147]
	_ = x[PruneSelectCols-148]
	_ = x[PruneLimitCols-149]
	_ = x[PruneOffsetCols-150]
	_ = x[PruneJoinLeftCols-151]
	_ = x[PruneJoinRightCols-152]
	_ = x[PruneSemiAntiJoinRightCols-153]
	_ = x[PruneAggCols-154]
	_ = x[PruneGroupByCols-155]
	_ = x[PruneValuesCols-156]
	_ = x[PruneOrdinalityCols-157]
	_ = x[PruneExplainCols-158]
	_ = x[PruneProjectSetCols-159]
	_ = x[PruneWindowOutputCols-160]
	_ = x[PruneWindowInputCols-161]
	_ = x[PruneMutationFetchCols-162]
	_ = x[PruneMutationInputCols-163]
	_ = x[PruneMutationReturnCols-164]
	_ = x[PruneWithScanCols-165]
	_ = x[PruneWithCols-166]
	_ = x[PruneUnionAllCols-167]
	_ = x[RejectNullsLeftJoin-168]
	_ = x[RejectNullsRightJoin-169]
	_ = x[RejectNullsGroupBy-170]
	_ = x[CommuteVar-171]
	_ = x[CommuteConst-172]
	_ = x[EliminateCoalesce-173]
	_ = x[SimplifyCoalesce-174]
	_ = x[EliminateCast-175]
	_ = x[NormalizeInConst-176]
	_ = x[FoldInNull-177]
	_ = x[UnifyComparisonTypes-178]
	_ = x[EliminateExistsProject-179]
	_ = x[EliminateExistsGroupBy-180]
	_ = x[IntroduceExistsLimit-181]
	_ = x[EliminateExistsLimit-182]
	_ = x[NormalizeJSONFieldAccess-183]
	_ = x[NormalizeJSONContains-184]
	_ = x[SimplifyCaseWhenConstValue-185]
	_ = x[InlineAnyValuesSingleCol-186]
	_ = x[InlineAnyValuesMultiCol-187]
	_ = x[SimplifyEqualsAnyTuple-188]
	_ = x[SimplifyAnyScalarArray-189]
	_ = x[FoldCollate-190]
	_ = x[NormalizeArrayFlattenToAgg-191]
	_ = x[SimplifySelectFilters-192]
	_ = x[ConsolidateSelectFilters-193]
	_ = x[DetectSelectContradiction-194]
	_ = x[EliminateSelect-195]
	_ = x[MergeSelects-196]
	_ = x[PushSelectIntoProject-197]
	_ = x[MergeSelectInnerJoin-198]
	_ = x[PushSelectCondLeftIntoJoinLeftAndRight-199]
	_ = x[PushSelectIntoJoinLeft-200]
	_ = x[PushSelectIntoGroupBy-201]
	_ = x[RemoveNotNullCondition-202]
	_ = x[InlineConstVar-203]
	_ = x[PushSelectIntoProjectSet-204]
	_ = x[PushFilterIntoSetOp-205]
	_ = x[EliminateUnionAllLeft-206]
	_ = x[EliminateUnionAllRight-207]
	_ = x[EliminateWindow-208]
	_ = x[ReduceWindowPartitionCols-209]
	_ = x[SimplifyWindowOrdering-210]
	_ = x[PushSelectIntoWindow-211]
	_ = x[PushLimitIntoWindow-212]
	_ = x[InlineWith-213]
	_ = x[startExploreRule-214]
	_ = x[ReplaceScalarMinMaxWithLimit-215]
	_ = x[ReplaceMinWithLimit-216]
	_ = x[ReplaceMaxWithLimit-217]
	_ = x[GenerateStreamingGroupBy-218]
	_ = x[CommuteJoin-219]
	_ = x[CommuteLeftJoin-220]
	_ = x[CommuteSemiJoin-221]
	_ = x[GenerateMergeJoins-222]
	_ = x[GenerateLookupJoins-223]
	_ = x[GenerateZigzagJoins-224]
	_ = x[GenerateInvertedIndexZigzagJoins-225]
	_ = x[GenerateLookupJoinsWithFilter-226]
	_ = x[AssociateJoin-227]
	_ = x[GenerateLimitedScans-228]
	_ = x[PushLimitIntoConstrainedScan-229]
	_ = x[PushLimitIntoIndexJoin-230]
	_ = x[GenerateIndexScans-231]
	_ = x[GenerateConstrainedScans-232]
	_ = x[GenerateInvertedIndexScans-233]
	_ = x[SplitDisjunction-234]
	_ = x[SplitDisjunctionAddKey-235]
	_ = x[NumRuleNames-236]
}

const _RuleName_name = "InvalidRuleNameSimplifyRootOrderingPruneRootColsSimplifyZeroCardinalityGroupNumManualRuleNamesEliminateAggDistinctNormalizeNestedAndsSimplifyTrueAndSimplifyAndTrueSimplifyFalseAndSimplifyAndFalseSimplifyTrueOrSimplifyOrTrueSimplifyFalseOrSimplifyOrFalseSimplifyRangeFoldNullAndOrFoldNotTrueFoldNotFalseFoldNotNullNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantConjunctCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldIsNotNullFoldNonNullIsNotNullCommuteNullIsDecorrelateJoinDecorrelateProjectSetTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateProjectInnerJoinTryDecorrelateInnerJoinTryDecorrelateInnerLeftJoinTryDecorrelateGroupByTryDecorrelateScalarGroupByTryDecorrelateSemiJoinTryDecorrelateLimitOneTryDecorrelateProjectSetTryDecorrelateWindowHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryHoistProjectSetSubqueryNormalizeSelectAnyFilterNormalizeJoinAnyFilterNormalizeSelectNotAnyFilterNormalizeJoinNotAnyFilterFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldInEmptyFoldNotInEmptyFoldArrayFoldBinaryFoldUnaryFoldComparisonFoldCastFoldIndirectionFoldColumnAccessFoldFunctionFoldEqualsAnyNullConvertGroupByToDistinctEliminateDistinctEliminateGroupByProjectReduceGroupingColsReduceNotNullGroupingColsEliminateAggDistinctForKeysEliminateAggFilteredDistinctForKeysEliminateDistinctNoColumnsEliminateErrorDistinctNoColumnsEliminateDistinctOnValuesPushAggDistinctIntoScalarGroupByPushAggFilterIntoScalarGroupByConvertCountToCountRowsInlineProjectConstantsInlineSelectConstantsInlineJoinConstantsLeftInlineJoinConstantsRightPushSelectIntoInlinableProjectInlineProjectInProjectCommuteRightJoinSimplifyJoinFiltersDetectJoinContradictionPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightMapEqualityIntoJoinLeftAndRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinWithoutFiltersSimplifyRightJoinWithoutFiltersSimplifyLeftJoinWithFiltersSimplifyRightJoinWithFiltersEliminateSemiJoinSimplifyZeroCardinalitySemiJoinEliminateAntiJoinSimplifyZeroCardinalityAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightHoistJoinProjectRightHoistJoinProjectLeftSimplifyJoinNotNullEqualityExtractJoinEqualitiesSortFiltersInJoinEliminateLimitEliminateOffsetPushLimitIntoProjectPushOffsetIntoProjectPushLimitIntoOffsetPushLimitIntoOrdinalityPushLimitIntoLeftJoinEliminateMax1RowFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyOrdinalityOrderingSimplifyExplainOrderingEliminateProjectMergeProjectsMergeProjectWithValuesPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneSemiAntiJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneOrdinalityColsPruneExplainColsPruneProjectSetColsPruneWindowOutputColsPruneWindowInputColsPruneMutationFetchColsPruneMutationInputColsPruneMutationReturnColsPruneWithScanColsPruneWithColsPruneUnionAllColsRejectNullsLeftJoinRejectNullsRightJoinRejectNullsGroupByCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastNormalizeInConstFoldInNullUnifyComparisonTypesEliminateExistsProjectEliminateExistsGroupByIntroduceExistsLimitEliminateExistsLimitNormalizeJSONFieldAccessNormalizeJSONContainsSimplifyCaseWhenConstValueInlineAnyValuesSingleColInlineAnyValuesMultiColSimplifyEqualsAnyTupleSimplifyAnyScalarArrayFoldCollateNormalizeArrayFlattenToAggSimplifySelectFiltersConsolidateSelectFiltersDetectSelectContradictionEliminateSelectMergeSelectsPushSelectIntoProjectMergeSelectInnerJoinPushSelectCondLeftIntoJoinLeftAndRightPushSelectIntoJoinLeftPushSelectIntoGroupByRemoveNotNullConditionInlineConstVarPushSelectIntoProjectSetPushFilterIntoSetOpEliminateUnionAllLeftEliminateUnionAllRightEliminateWindowReduceWindowPartitionColsSimplifyWindowOrderingPushSelectIntoWindowPushLimitIntoWindowInlineWithstartExploreRuleReplaceScalarMinMaxWithLimitReplaceMinWithLimitReplaceMaxWithLimitGenerateStreamingGroupByCommuteJoinCommuteLeftJoinCommuteSemiJoinGenerateMergeJoinsGenerateLookupJoinsGenerateZigzagJoinsGenerateInvertedIndexZigzagJoinsGenerateLookupJoinsWithFilterAssociateJoinGenerateLimitedScansPushLimitIntoConstrainedScanPushLimitIntoIndexJoinGenerateIndexScansGenerateConstrainedScansGenerateInvertedIndexScansSplitDisjunctionSplitDisjunctionAddKeyNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 35, 48, 76, 94, 114, 133, 148, 163, 179, 195, 209, 223, 238, 253, 266, 279, 290, 302, 313, 329, 341, 350, 358, 382, 402, 424, 445, 467, 489, 511, 533, 556, 566, 583, 596, 616, 629, 644, 665, 685, 706, 733, 763, 786, 813, 834, 861, 883, 905, 929, 949, 966, 986, 1005, 1025, 1042, 1061, 1084, 1108, 1130, 1157, 1182, 1194, 1207, 1225, 1244, 1262, 1273, 1287, 1296, 1306, 1315, 1329, 1337, 1352, 1368, 1380, 1397, 1421, 1438, 1461, 1479, 1504, 1531, 1566, 1592, 1623, 1648, 1680, 1710, 1733, 1755, 1776, 1799, 1823, 1853, 1875, 1891, 1910, 1933, 1963, 1984, 2006, 2037, 2059, 2082, 2112, 2143, 2170, 2198, 2215, 2246, 2263, 2294, 2317, 2341, 2362, 2382, 2409, 2430, 2447, 2461, 2476, 2496, 2517, 2536, 2559, 2580, 2596, 2608, 2620, 2633, 2644, 2655, 2665, 2676, 2695, 2716, 2738, 2761, 2787, 2810, 2826, 2839, 2861, 2877, 2890, 2905, 2919, 2934, 2951, 2969, 2995, 3007, 3023, 3038, 3057, 3073, 3092, 3113, 3133, 3155, 3177, 3200, 3217, 3230, 3247, 3266, 3286, 3304, 3314, 3326, 3343, 3359, 3372, 3388, 3398, 3418, 3440, 3462, 3482, 3502, 3526, 3547, 3573, 3597, 3620, 3642, 3664, 3675, 3701, 3722, 3746, 3771, 3786, 3798, 3819, 3839, 3877, 3899, 3920, 3942, 3956, 3980, 3999, 4020, 4042, 4057, 4082, 4104, 4124, 4143, 4153, 4169, 4197, 4216, 4235, 4259, 4270, 4285, 4300, 4318, 4337, 4356, 4388, 4417, 4430, 4450, 4478, 4500, 4518, 4542, 4568, 4584, 4606, 4618}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
