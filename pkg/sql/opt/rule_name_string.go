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
	_ = x[FoldNullTupleIsTupleNull-35]
	_ = x[FoldNonNullTupleIsTupleNull-36]
	_ = x[FoldIsNotNull-37]
	_ = x[FoldNonNullIsNotNull-38]
	_ = x[FoldNonNullTupleIsTupleNotNull-39]
	_ = x[FoldNullTupleIsTupleNotNull-40]
	_ = x[CommuteNullIs-41]
	_ = x[DecorrelateJoin-42]
	_ = x[DecorrelateProjectSet-43]
	_ = x[TryDecorrelateSelect-44]
	_ = x[TryDecorrelateProject-45]
	_ = x[TryDecorrelateProjectSelect-46]
	_ = x[TryDecorrelateProjectInnerJoin-47]
	_ = x[TryDecorrelateInnerJoin-48]
	_ = x[TryDecorrelateInnerLeftJoin-49]
	_ = x[TryDecorrelateGroupBy-50]
	_ = x[TryDecorrelateScalarGroupBy-51]
	_ = x[TryDecorrelateSemiJoin-52]
	_ = x[TryDecorrelateLimitOne-53]
	_ = x[TryDecorrelateProjectSet-54]
	_ = x[TryDecorrelateWindow-55]
	_ = x[TryDecorrelateMax1Row-56]
	_ = x[HoistSelectExists-57]
	_ = x[HoistSelectNotExists-58]
	_ = x[HoistSelectSubquery-59]
	_ = x[HoistProjectSubquery-60]
	_ = x[HoistJoinSubquery-61]
	_ = x[HoistValuesSubquery-62]
	_ = x[HoistProjectSetSubquery-63]
	_ = x[NormalizeSelectAnyFilter-64]
	_ = x[NormalizeJoinAnyFilter-65]
	_ = x[NormalizeSelectNotAnyFilter-66]
	_ = x[NormalizeJoinNotAnyFilter-67]
	_ = x[FoldNullCast-68]
	_ = x[FoldNullUnary-69]
	_ = x[FoldNullBinaryLeft-70]
	_ = x[FoldNullBinaryRight-71]
	_ = x[FoldNullInNonEmpty-72]
	_ = x[FoldInEmpty-73]
	_ = x[FoldNotInEmpty-74]
	_ = x[FoldArray-75]
	_ = x[FoldBinary-76]
	_ = x[FoldUnary-77]
	_ = x[FoldComparison-78]
	_ = x[FoldCast-79]
	_ = x[FoldIndirection-80]
	_ = x[FoldColumnAccess-81]
	_ = x[FoldFunction-82]
	_ = x[FoldEqualsAnyNull-83]
	_ = x[ConvertGroupByToDistinct-84]
	_ = x[EliminateDistinct-85]
	_ = x[EliminateGroupByProject-86]
	_ = x[ReduceGroupingCols-87]
	_ = x[ReduceNotNullGroupingCols-88]
	_ = x[EliminateAggDistinctForKeys-89]
	_ = x[EliminateAggFilteredDistinctForKeys-90]
	_ = x[EliminateDistinctNoColumns-91]
	_ = x[EliminateEnsureDistinctNoColumns-92]
	_ = x[EliminateDistinctOnValues-93]
	_ = x[PushAggDistinctIntoScalarGroupBy-94]
	_ = x[PushAggFilterIntoScalarGroupBy-95]
	_ = x[ConvertCountToCountRows-96]
	_ = x[InlineConstVar-97]
	_ = x[InlineProjectConstants-98]
	_ = x[InlineSelectConstants-99]
	_ = x[InlineJoinConstantsLeft-100]
	_ = x[InlineJoinConstantsRight-101]
	_ = x[PushSelectIntoInlinableProject-102]
	_ = x[InlineProjectInProject-103]
	_ = x[CommuteRightJoin-104]
	_ = x[SimplifyJoinFilters-105]
	_ = x[DetectJoinContradiction-106]
	_ = x[PushFilterIntoJoinLeftAndRight-107]
	_ = x[MapFilterIntoJoinLeft-108]
	_ = x[MapFilterIntoJoinRight-109]
	_ = x[MapEqualityIntoJoinLeftAndRight-110]
	_ = x[PushFilterIntoJoinLeft-111]
	_ = x[PushFilterIntoJoinRight-112]
	_ = x[SimplifyLeftJoinWithoutFilters-113]
	_ = x[SimplifyRightJoinWithoutFilters-114]
	_ = x[SimplifyLeftJoinWithFilters-115]
	_ = x[SimplifyRightJoinWithFilters-116]
	_ = x[EliminateSemiJoin-117]
	_ = x[SimplifyZeroCardinalitySemiJoin-118]
	_ = x[EliminateAntiJoin-119]
	_ = x[SimplifyZeroCardinalityAntiJoin-120]
	_ = x[EliminateJoinNoColsLeft-121]
	_ = x[EliminateJoinNoColsRight-122]
	_ = x[HoistJoinProjectRight-123]
	_ = x[HoistJoinProjectLeft-124]
	_ = x[SimplifyJoinNotNullEquality-125]
	_ = x[ExtractJoinEqualities-126]
	_ = x[SortFiltersInJoin-127]
	_ = x[EliminateLimit-128]
	_ = x[EliminateOffset-129]
	_ = x[PushLimitIntoProject-130]
	_ = x[PushOffsetIntoProject-131]
	_ = x[PushLimitIntoOffset-132]
	_ = x[PushLimitIntoOrdinality-133]
	_ = x[PushLimitIntoLeftJoin-134]
	_ = x[EliminateMax1Row-135]
	_ = x[FoldPlusZero-136]
	_ = x[FoldZeroPlus-137]
	_ = x[FoldMinusZero-138]
	_ = x[FoldMultOne-139]
	_ = x[FoldOneMult-140]
	_ = x[FoldDivOne-141]
	_ = x[InvertMinus-142]
	_ = x[EliminateUnaryMinus-143]
	_ = x[SimplifyLimitOrdering-144]
	_ = x[SimplifyOffsetOrdering-145]
	_ = x[SimplifyGroupByOrdering-146]
	_ = x[SimplifyOrdinalityOrdering-147]
	_ = x[SimplifyExplainOrdering-148]
	_ = x[EliminateProject-149]
	_ = x[MergeProjects-150]
	_ = x[MergeProjectWithValues-151]
	_ = x[PushColumnRemappingIntoValues-152]
	_ = x[FoldTupleAccessIntoValues-153]
	_ = x[ConvertZipArraysToValues-154]
	_ = x[PruneProjectCols-155]
	_ = x[PruneScanCols-156]
	_ = x[PruneSelectCols-157]
	_ = x[PruneLimitCols-158]
	_ = x[PruneOffsetCols-159]
	_ = x[PruneJoinLeftCols-160]
	_ = x[PruneJoinRightCols-161]
	_ = x[PruneSemiAntiJoinRightCols-162]
	_ = x[PruneAggCols-163]
	_ = x[PruneGroupByCols-164]
	_ = x[PruneValuesCols-165]
	_ = x[PruneOrdinalityCols-166]
	_ = x[PruneExplainCols-167]
	_ = x[PruneProjectSetCols-168]
	_ = x[PruneWindowOutputCols-169]
	_ = x[PruneWindowInputCols-170]
	_ = x[PruneMutationFetchCols-171]
	_ = x[PruneMutationInputCols-172]
	_ = x[PruneMutationReturnCols-173]
	_ = x[PruneWithScanCols-174]
	_ = x[PruneWithCols-175]
	_ = x[PruneUnionAllCols-176]
	_ = x[RejectNullsLeftJoin-177]
	_ = x[RejectNullsRightJoin-178]
	_ = x[RejectNullsGroupBy-179]
	_ = x[CommuteVar-180]
	_ = x[CommuteConst-181]
	_ = x[EliminateCoalesce-182]
	_ = x[SimplifyCoalesce-183]
	_ = x[EliminateCast-184]
	_ = x[NormalizeInConst-185]
	_ = x[FoldInNull-186]
	_ = x[UnifyComparisonTypes-187]
	_ = x[EliminateExistsZeroRows-188]
	_ = x[EliminateExistsProject-189]
	_ = x[EliminateExistsGroupBy-190]
	_ = x[IntroduceExistsLimit-191]
	_ = x[EliminateExistsLimit-192]
	_ = x[NormalizeJSONFieldAccess-193]
	_ = x[NormalizeJSONContains-194]
	_ = x[SimplifyCaseWhenConstValue-195]
	_ = x[InlineAnyValuesSingleCol-196]
	_ = x[InlineAnyValuesMultiCol-197]
	_ = x[SimplifyEqualsAnyTuple-198]
	_ = x[SimplifyAnyScalarArray-199]
	_ = x[FoldCollate-200]
	_ = x[NormalizeArrayFlattenToAgg-201]
	_ = x[SimplifySameVarEqualities-202]
	_ = x[SimplifySameVarInequalities-203]
	_ = x[SimplifySelectFilters-204]
	_ = x[ConsolidateSelectFilters-205]
	_ = x[DetectSelectContradiction-206]
	_ = x[EliminateSelect-207]
	_ = x[MergeSelects-208]
	_ = x[PushSelectIntoProject-209]
	_ = x[MergeSelectInnerJoin-210]
	_ = x[PushSelectCondLeftIntoJoinLeftAndRight-211]
	_ = x[PushSelectIntoJoinLeft-212]
	_ = x[PushSelectIntoGroupBy-213]
	_ = x[RemoveNotNullCondition-214]
	_ = x[PushSelectIntoProjectSet-215]
	_ = x[PushFilterIntoSetOp-216]
	_ = x[EliminateUnionAllLeft-217]
	_ = x[EliminateUnionAllRight-218]
	_ = x[EliminateWindow-219]
	_ = x[ReduceWindowPartitionCols-220]
	_ = x[SimplifyWindowOrdering-221]
	_ = x[PushSelectIntoWindow-222]
	_ = x[PushLimitIntoWindow-223]
	_ = x[InlineWith-224]
	_ = x[startExploreRule-225]
	_ = x[ReplaceScalarMinMaxWithLimit-226]
	_ = x[ReplaceMinWithLimit-227]
	_ = x[ReplaceMaxWithLimit-228]
	_ = x[GenerateStreamingGroupBy-229]
	_ = x[CommuteJoin-230]
	_ = x[CommuteLeftJoin-231]
	_ = x[CommuteSemiJoin-232]
	_ = x[GenerateMergeJoins-233]
	_ = x[GenerateLookupJoins-234]
	_ = x[GenerateGeoLookupJoins-235]
	_ = x[GenerateZigzagJoins-236]
	_ = x[GenerateInvertedIndexZigzagJoins-237]
	_ = x[GenerateLookupJoinsWithFilter-238]
	_ = x[AssociateJoin-239]
	_ = x[GenerateLimitedScans-240]
	_ = x[PushLimitIntoConstrainedScan-241]
	_ = x[PushLimitIntoIndexJoin-242]
	_ = x[GenerateIndexScans-243]
	_ = x[GenerateConstrainedScans-244]
	_ = x[GenerateInvertedIndexScans-245]
	_ = x[SplitDisjunction-246]
	_ = x[SplitDisjunctionAddKey-247]
	_ = x[NumRuleNames-248]
}

const _RuleName_name = "InvalidRuleNameSimplifyRootOrderingPruneRootColsSimplifyZeroCardinalityGroupNumManualRuleNamesEliminateAggDistinctNormalizeNestedAndsSimplifyTrueAndSimplifyAndTrueSimplifyFalseAndSimplifyAndFalseSimplifyTrueOrSimplifyOrTrueSimplifyFalseOrSimplifyOrFalseSimplifyRangeFoldNullAndOrFoldNotTrueFoldNotFalseFoldNotNullNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantConjunctCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldNullTupleIsTupleNullFoldNonNullTupleIsTupleNullFoldIsNotNullFoldNonNullIsNotNullFoldNonNullTupleIsTupleNotNullFoldNullTupleIsTupleNotNullCommuteNullIsDecorrelateJoinDecorrelateProjectSetTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateProjectInnerJoinTryDecorrelateInnerJoinTryDecorrelateInnerLeftJoinTryDecorrelateGroupByTryDecorrelateScalarGroupByTryDecorrelateSemiJoinTryDecorrelateLimitOneTryDecorrelateProjectSetTryDecorrelateWindowTryDecorrelateMax1RowHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryHoistProjectSetSubqueryNormalizeSelectAnyFilterNormalizeJoinAnyFilterNormalizeSelectNotAnyFilterNormalizeJoinNotAnyFilterFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldInEmptyFoldNotInEmptyFoldArrayFoldBinaryFoldUnaryFoldComparisonFoldCastFoldIndirectionFoldColumnAccessFoldFunctionFoldEqualsAnyNullConvertGroupByToDistinctEliminateDistinctEliminateGroupByProjectReduceGroupingColsReduceNotNullGroupingColsEliminateAggDistinctForKeysEliminateAggFilteredDistinctForKeysEliminateDistinctNoColumnsEliminateEnsureDistinctNoColumnsEliminateDistinctOnValuesPushAggDistinctIntoScalarGroupByPushAggFilterIntoScalarGroupByConvertCountToCountRowsInlineConstVarInlineProjectConstantsInlineSelectConstantsInlineJoinConstantsLeftInlineJoinConstantsRightPushSelectIntoInlinableProjectInlineProjectInProjectCommuteRightJoinSimplifyJoinFiltersDetectJoinContradictionPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightMapEqualityIntoJoinLeftAndRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinWithoutFiltersSimplifyRightJoinWithoutFiltersSimplifyLeftJoinWithFiltersSimplifyRightJoinWithFiltersEliminateSemiJoinSimplifyZeroCardinalitySemiJoinEliminateAntiJoinSimplifyZeroCardinalityAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightHoistJoinProjectRightHoistJoinProjectLeftSimplifyJoinNotNullEqualityExtractJoinEqualitiesSortFiltersInJoinEliminateLimitEliminateOffsetPushLimitIntoProjectPushOffsetIntoProjectPushLimitIntoOffsetPushLimitIntoOrdinalityPushLimitIntoLeftJoinEliminateMax1RowFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyOrdinalityOrderingSimplifyExplainOrderingEliminateProjectMergeProjectsMergeProjectWithValuesPushColumnRemappingIntoValuesFoldTupleAccessIntoValuesConvertZipArraysToValuesPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneSemiAntiJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneOrdinalityColsPruneExplainColsPruneProjectSetColsPruneWindowOutputColsPruneWindowInputColsPruneMutationFetchColsPruneMutationInputColsPruneMutationReturnColsPruneWithScanColsPruneWithColsPruneUnionAllColsRejectNullsLeftJoinRejectNullsRightJoinRejectNullsGroupByCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastNormalizeInConstFoldInNullUnifyComparisonTypesEliminateExistsZeroRowsEliminateExistsProjectEliminateExistsGroupByIntroduceExistsLimitEliminateExistsLimitNormalizeJSONFieldAccessNormalizeJSONContainsSimplifyCaseWhenConstValueInlineAnyValuesSingleColInlineAnyValuesMultiColSimplifyEqualsAnyTupleSimplifyAnyScalarArrayFoldCollateNormalizeArrayFlattenToAggSimplifySameVarEqualitiesSimplifySameVarInequalitiesSimplifySelectFiltersConsolidateSelectFiltersDetectSelectContradictionEliminateSelectMergeSelectsPushSelectIntoProjectMergeSelectInnerJoinPushSelectCondLeftIntoJoinLeftAndRightPushSelectIntoJoinLeftPushSelectIntoGroupByRemoveNotNullConditionPushSelectIntoProjectSetPushFilterIntoSetOpEliminateUnionAllLeftEliminateUnionAllRightEliminateWindowReduceWindowPartitionColsSimplifyWindowOrderingPushSelectIntoWindowPushLimitIntoWindowInlineWithstartExploreRuleReplaceScalarMinMaxWithLimitReplaceMinWithLimitReplaceMaxWithLimitGenerateStreamingGroupByCommuteJoinCommuteLeftJoinCommuteSemiJoinGenerateMergeJoinsGenerateLookupJoinsGenerateGeoLookupJoinsGenerateZigzagJoinsGenerateInvertedIndexZigzagJoinsGenerateLookupJoinsWithFilterAssociateJoinGenerateLimitedScansPushLimitIntoConstrainedScanPushLimitIntoIndexJoinGenerateIndexScansGenerateConstrainedScansGenerateInvertedIndexScansSplitDisjunctionSplitDisjunctionAddKeyNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 35, 48, 76, 94, 114, 133, 148, 163, 179, 195, 209, 223, 238, 253, 266, 279, 290, 302, 313, 329, 341, 350, 358, 382, 402, 424, 445, 467, 489, 511, 533, 556, 566, 583, 607, 634, 647, 667, 697, 724, 737, 752, 773, 793, 814, 841, 871, 894, 921, 942, 969, 991, 1013, 1037, 1057, 1078, 1095, 1115, 1134, 1154, 1171, 1190, 1213, 1237, 1259, 1286, 1311, 1323, 1336, 1354, 1373, 1391, 1402, 1416, 1425, 1435, 1444, 1458, 1466, 1481, 1497, 1509, 1526, 1550, 1567, 1590, 1608, 1633, 1660, 1695, 1721, 1753, 1778, 1810, 1840, 1863, 1877, 1899, 1920, 1943, 1967, 1997, 2019, 2035, 2054, 2077, 2107, 2128, 2150, 2181, 2203, 2226, 2256, 2287, 2314, 2342, 2359, 2390, 2407, 2438, 2461, 2485, 2506, 2526, 2553, 2574, 2591, 2605, 2620, 2640, 2661, 2680, 2703, 2724, 2740, 2752, 2764, 2777, 2788, 2799, 2809, 2820, 2839, 2860, 2882, 2905, 2931, 2954, 2970, 2983, 3005, 3034, 3059, 3083, 3099, 3112, 3127, 3141, 3156, 3173, 3191, 3217, 3229, 3245, 3260, 3279, 3295, 3314, 3335, 3355, 3377, 3399, 3422, 3439, 3452, 3469, 3488, 3508, 3526, 3536, 3548, 3565, 3581, 3594, 3610, 3620, 3640, 3663, 3685, 3707, 3727, 3747, 3771, 3792, 3818, 3842, 3865, 3887, 3909, 3920, 3946, 3971, 3998, 4019, 4043, 4068, 4083, 4095, 4116, 4136, 4174, 4196, 4217, 4239, 4263, 4282, 4303, 4325, 4340, 4365, 4387, 4407, 4426, 4436, 4452, 4480, 4499, 4518, 4542, 4553, 4568, 4583, 4601, 4620, 4642, 4661, 4693, 4722, 4735, 4755, 4783, 4805, 4823, 4847, 4873, 4889, 4911, 4923}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
