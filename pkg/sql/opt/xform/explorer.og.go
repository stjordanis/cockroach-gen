// Code generated by optgen; DO NOT EDIT.

package xform

import (
	"github.com/cockroachdb/cockroach/pkg/sql/opt"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/memo"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

func (_e *explorer) exploreGroupMember(
	state *exploreState,
	member memo.RelExpr,
	ordinal int,
) (_fullyExplored bool) {
	switch t := member.(type) {
	case *memo.ScanExpr:
		return _e.exploreScan(state, t, ordinal)
	case *memo.SelectExpr:
		return _e.exploreSelect(state, t, ordinal)
	case *memo.InnerJoinExpr:
		return _e.exploreInnerJoin(state, t, ordinal)
	case *memo.LeftJoinExpr:
		return _e.exploreLeftJoin(state, t, ordinal)
	case *memo.RightJoinExpr:
		return _e.exploreRightJoin(state, t, ordinal)
	case *memo.FullJoinExpr:
		return _e.exploreFullJoin(state, t, ordinal)
	case *memo.SemiJoinExpr:
		return _e.exploreSemiJoin(state, t, ordinal)
	case *memo.AntiJoinExpr:
		return _e.exploreAntiJoin(state, t, ordinal)
	case *memo.GroupByExpr:
		return _e.exploreGroupBy(state, t, ordinal)
	case *memo.ScalarGroupByExpr:
		return _e.exploreScalarGroupBy(state, t, ordinal)
	case *memo.DistinctOnExpr:
		return _e.exploreDistinctOn(state, t, ordinal)
	case *memo.LimitExpr:
		return _e.exploreLimit(state, t, ordinal)
	}

	// No rules for other operator types.
	return true
}

func (_e *explorer) exploreScan(
	_rootState *exploreState,
	_root *memo.ScanExpr,
	_rootOrd int,
) (_fullyExplored bool) {
	_fullyExplored = true

	// [GenerateIndexScans]
	{
		if _rootOrd >= _rootState.start {
			scanPrivate := &_root.ScanPrivate
			if _e.funcs.IsCanonicalScan(scanPrivate) {
				if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateIndexScans) {
					var _last memo.RelExpr
					if _e.o.appliedRule != nil {
						_last = memo.LastGroupMember(_root)
					}
					_e.funcs.GenerateIndexScans(_root, scanPrivate)
					if _e.o.appliedRule != nil {
						_e.o.appliedRule(opt.GenerateIndexScans, _root, _last.NextExpr())
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreSelect(
	_rootState *exploreState,
	_root *memo.SelectExpr,
	_rootOrd int,
) (_fullyExplored bool) {
	_fullyExplored = true

	// [GenerateZigzagJoins]
	{
		_partlyExplored := _rootOrd < _rootState.start
		_state := _e.lookupExploreState(_root.Input)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = _root.Input.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			if !_partlyExplored || _ord >= _state.start {
				_scan, _ := _member.(*memo.ScanExpr)
				if _scan != nil {
					scan := &_scan.ScanPrivate
					if _e.funcs.IsCanonicalScan(scan) {
						filters := _root.Filters
						if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateZigzagJoins) {
							var _last memo.RelExpr
							if _e.o.appliedRule != nil {
								_last = memo.LastGroupMember(_root)
							}
							_e.funcs.GenerateZigzagJoins(_root, scan, filters)
							if _e.o.appliedRule != nil {
								_e.o.appliedRule(opt.GenerateZigzagJoins, _root, _last.NextExpr())
							}
						}
					}
				}
			}
		}
	}

	// [GenerateInvertedIndexZigzagJoins]
	{
		_partlyExplored := _rootOrd < _rootState.start
		_state := _e.lookupExploreState(_root.Input)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = _root.Input.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			if !_partlyExplored || _ord >= _state.start {
				_scan, _ := _member.(*memo.ScanExpr)
				if _scan != nil {
					scan := &_scan.ScanPrivate
					if _e.funcs.IsCanonicalScan(scan) {
						if _e.funcs.HasInvertedIndexes(scan) {
							filters := _root.Filters
							if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateInvertedIndexZigzagJoins) {
								var _last memo.RelExpr
								if _e.o.appliedRule != nil {
									_last = memo.LastGroupMember(_root)
								}
								_e.funcs.GenerateInvertedIndexZigzagJoins(_root, scan, filters)
								if _e.o.appliedRule != nil {
									_e.o.appliedRule(opt.GenerateInvertedIndexZigzagJoins, _root, _last.NextExpr())
								}
							}
						}
					}
				}
			}
		}
	}

	// [GenerateConstrainedScans]
	{
		_partlyExplored := _rootOrd < _rootState.start
		_state := _e.lookupExploreState(_root.Input)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = _root.Input.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			if !_partlyExplored || _ord >= _state.start {
				_scan, _ := _member.(*memo.ScanExpr)
				if _scan != nil {
					scanPrivate := &_scan.ScanPrivate
					if _e.funcs.IsCanonicalScan(scanPrivate) {
						filters := _root.Filters
						if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateConstrainedScans) {
							var _last memo.RelExpr
							if _e.o.appliedRule != nil {
								_last = memo.LastGroupMember(_root)
							}
							_e.funcs.GenerateConstrainedScans(_root, scanPrivate, filters)
							if _e.o.appliedRule != nil {
								_e.o.appliedRule(opt.GenerateConstrainedScans, _root, _last.NextExpr())
							}
						}
					}
				}
			}
		}
	}

	// [GenerateInvertedIndexScans]
	{
		_partlyExplored := _rootOrd < _rootState.start
		_state := _e.lookupExploreState(_root.Input)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = _root.Input.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			if !_partlyExplored || _ord >= _state.start {
				_scan, _ := _member.(*memo.ScanExpr)
				if _scan != nil {
					scanPrivate := &_scan.ScanPrivate
					if _e.funcs.IsCanonicalScan(scanPrivate) {
						if _e.funcs.HasInvertedIndexes(scanPrivate) {
							filters := _root.Filters
							if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateInvertedIndexScans) {
								var _last memo.RelExpr
								if _e.o.appliedRule != nil {
									_last = memo.LastGroupMember(_root)
								}
								_e.funcs.GenerateInvertedIndexScans(_root, scanPrivate, filters)
								if _e.o.appliedRule != nil {
									_e.o.appliedRule(opt.GenerateInvertedIndexScans, _root, _last.NextExpr())
								}
							}
						}
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreInnerJoin(
	_rootState *exploreState,
	_root *memo.InnerJoinExpr,
	_rootOrd int,
) (_fullyExplored bool) {
	_fullyExplored = true

	// [CommuteJoin]
	{
		if _rootOrd >= _rootState.start {
			left := _root.Left
			right := _root.Right
			on := _root.On
			private := &_root.JoinPrivate
			if _e.funcs.NoJoinHints(private) {
				if _e.o.matchedRule == nil || _e.o.matchedRule(opt.CommuteJoin) {
					_expr := &memo.InnerJoinExpr{
						Left:        right,
						Right:       left,
						On:          on,
						JoinPrivate: *private,
					}
					_interned := _e.mem.AddInnerJoinToGroup(_expr, _root)
					if _e.o.appliedRule != nil {
						if _interned != _expr {
							_e.o.appliedRule(opt.CommuteJoin, _root, nil)
						} else {
							_e.o.appliedRule(opt.CommuteJoin, _root, _interned)
						}
					}
				}
			}
		}
	}

	// [GenerateMergeJoins]
	{
		if _rootOrd >= _rootState.start {
			left := _root.Left
			right := _root.Right
			on := _root.On
			private := &_root.JoinPrivate
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateMergeJoins) {
				var _last memo.RelExpr
				if _e.o.appliedRule != nil {
					_last = memo.LastGroupMember(_root)
				}
				_e.funcs.GenerateMergeJoins(_root, opt.InnerJoinOp, left, right, on, private)
				if _e.o.appliedRule != nil {
					_e.o.appliedRule(opt.GenerateMergeJoins, _root, _last.NextExpr())
				}
			}
		}
	}

	// [GenerateLookupJoins]
	{
		_partlyExplored := _rootOrd < _rootState.start
		left := _root.Left
		_state := _e.lookupExploreState(_root.Right)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = _root.Right.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			if !_partlyExplored || _ord >= _state.start {
				_scan, _ := _member.(*memo.ScanExpr)
				if _scan != nil {
					scanPrivate := &_scan.ScanPrivate
					if _e.funcs.IsCanonicalScan(scanPrivate) {
						on := _root.On
						private := &_root.JoinPrivate
						if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateLookupJoins) {
							var _last memo.RelExpr
							if _e.o.appliedRule != nil {
								_last = memo.LastGroupMember(_root)
							}
							_e.funcs.GenerateLookupJoins(_root, opt.InnerJoinOp, left, scanPrivate, on, private)
							if _e.o.appliedRule != nil {
								_e.o.appliedRule(opt.GenerateLookupJoins, _root, _last.NextExpr())
							}
						}
					}
				}
			}
		}
	}

	// [GenerateLookupJoinsWithFilter]
	{
		_partlyExplored := _rootOrd < _rootState.start
		left := _root.Left
		_state := _e.lookupExploreState(_root.Right)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = _root.Right.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			_partlyExplored := _partlyExplored && _ord < _state.start
			_select, _ := _member.(*memo.SelectExpr)
			if _select != nil {
				_state := _e.lookupExploreState(_select.Input)
				if !_state.fullyExplored {
					_fullyExplored = false
				}
				var _member memo.RelExpr
				for _ord := 0; _ord < _state.end; _ord++ {
					if _member == nil {
						_member = _select.Input.FirstExpr()
					} else {
						_member = _member.NextExpr()
					}
					if !_partlyExplored || _ord >= _state.start {
						_scan, _ := _member.(*memo.ScanExpr)
						if _scan != nil {
							scanPrivate := &_scan.ScanPrivate
							if _e.funcs.IsCanonicalScan(scanPrivate) {
								filters := _select.Filters
								on := _root.On
								private := &_root.JoinPrivate
								if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateLookupJoinsWithFilter) {
									var _last memo.RelExpr
									if _e.o.appliedRule != nil {
										_last = memo.LastGroupMember(_root)
									}
									_e.funcs.GenerateLookupJoins(_root, opt.InnerJoinOp, left, scanPrivate, _e.funcs.ConcatFilters(on, filters), private)
									if _e.o.appliedRule != nil {
										_e.o.appliedRule(opt.GenerateLookupJoinsWithFilter, _root, _last.NextExpr())
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// [AssociateJoin]
	{
		_partlyExplored := _rootOrd < _rootState.start
		left := _root.Left
		_state := _e.lookupExploreState(left)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = left.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			if !_partlyExplored || _ord >= _state.start {
				_innerJoin, _ := _member.(*memo.InnerJoinExpr)
				if _innerJoin != nil {
					innerLeft := _innerJoin.Left
					innerRight := _innerJoin.Right
					innerOn := _innerJoin.On
					innerPrivate := &_innerJoin.JoinPrivate
					if _e.funcs.NoJoinHints(innerPrivate) {
						right := _root.Right
						if _e.funcs.ShouldReorderJoins(left, right) {
							on := _root.On
							private := &_root.JoinPrivate
							if _e.funcs.NoJoinHints(private) {
								if _e.o.matchedRule == nil || _e.o.matchedRule(opt.AssociateJoin) {
									cols := _e.funcs.OutputCols2(innerRight, right)
									newOn := _e.funcs.MapJoinOpEqualities(_e.funcs.ConcatFilters(on, innerOn), _e.funcs.OutputCols(innerLeft), cols)
									_expr := &memo.InnerJoinExpr{
										Left: innerLeft,
										Right: _e.f.ConstructInnerJoin(
											innerRight,
											right,
											_e.funcs.SortFilters(_e.funcs.ExtractBoundConditions(newOn, cols)),
											_e.funcs.EmptyJoinPrivate(),
										),
										On:          _e.funcs.SortFilters(_e.funcs.ExtractUnboundConditions(newOn, cols)),
										JoinPrivate: *_e.funcs.EmptyJoinPrivate(),
									}
									_interned := _e.mem.AddInnerJoinToGroup(_expr, _root)
									if _e.o.appliedRule != nil {
										if _interned != _expr {
											_e.o.appliedRule(opt.AssociateJoin, _root, nil)
										} else {
											_e.o.appliedRule(opt.AssociateJoin, _root, _interned)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreLeftJoin(
	_rootState *exploreState,
	_root *memo.LeftJoinExpr,
	_rootOrd int,
) (_fullyExplored bool) {
	_fullyExplored = true

	// [CommuteLeftJoin]
	{
		if _rootOrd >= _rootState.start {
			left := _root.Left
			right := _root.Right
			on := _root.On
			private := &_root.JoinPrivate
			if _e.funcs.NoJoinHints(private) {
				if _e.o.matchedRule == nil || _e.o.matchedRule(opt.CommuteLeftJoin) {
					_expr := &memo.RightJoinExpr{
						Left:        right,
						Right:       left,
						On:          on,
						JoinPrivate: *private,
					}
					_interned := _e.mem.AddRightJoinToGroup(_expr, _root)
					if _e.o.appliedRule != nil {
						if _interned != _expr {
							_e.o.appliedRule(opt.CommuteLeftJoin, _root, nil)
						} else {
							_e.o.appliedRule(opt.CommuteLeftJoin, _root, _interned)
						}
					}
				}
			}
		}
	}

	// [GenerateMergeJoins]
	{
		if _rootOrd >= _rootState.start {
			left := _root.Left
			right := _root.Right
			on := _root.On
			private := &_root.JoinPrivate
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateMergeJoins) {
				var _last memo.RelExpr
				if _e.o.appliedRule != nil {
					_last = memo.LastGroupMember(_root)
				}
				_e.funcs.GenerateMergeJoins(_root, opt.LeftJoinOp, left, right, on, private)
				if _e.o.appliedRule != nil {
					_e.o.appliedRule(opt.GenerateMergeJoins, _root, _last.NextExpr())
				}
			}
		}
	}

	// [GenerateLookupJoins]
	{
		_partlyExplored := _rootOrd < _rootState.start
		left := _root.Left
		_state := _e.lookupExploreState(_root.Right)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = _root.Right.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			if !_partlyExplored || _ord >= _state.start {
				_scan, _ := _member.(*memo.ScanExpr)
				if _scan != nil {
					scanPrivate := &_scan.ScanPrivate
					if _e.funcs.IsCanonicalScan(scanPrivate) {
						on := _root.On
						private := &_root.JoinPrivate
						if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateLookupJoins) {
							var _last memo.RelExpr
							if _e.o.appliedRule != nil {
								_last = memo.LastGroupMember(_root)
							}
							_e.funcs.GenerateLookupJoins(_root, opt.LeftJoinOp, left, scanPrivate, on, private)
							if _e.o.appliedRule != nil {
								_e.o.appliedRule(opt.GenerateLookupJoins, _root, _last.NextExpr())
							}
						}
					}
				}
			}
		}
	}

	// [GenerateLookupJoinsWithFilter]
	{
		_partlyExplored := _rootOrd < _rootState.start
		left := _root.Left
		_state := _e.lookupExploreState(_root.Right)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = _root.Right.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			_partlyExplored := _partlyExplored && _ord < _state.start
			_select, _ := _member.(*memo.SelectExpr)
			if _select != nil {
				_state := _e.lookupExploreState(_select.Input)
				if !_state.fullyExplored {
					_fullyExplored = false
				}
				var _member memo.RelExpr
				for _ord := 0; _ord < _state.end; _ord++ {
					if _member == nil {
						_member = _select.Input.FirstExpr()
					} else {
						_member = _member.NextExpr()
					}
					if !_partlyExplored || _ord >= _state.start {
						_scan, _ := _member.(*memo.ScanExpr)
						if _scan != nil {
							scanPrivate := &_scan.ScanPrivate
							if _e.funcs.IsCanonicalScan(scanPrivate) {
								filters := _select.Filters
								on := _root.On
								private := &_root.JoinPrivate
								if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateLookupJoinsWithFilter) {
									var _last memo.RelExpr
									if _e.o.appliedRule != nil {
										_last = memo.LastGroupMember(_root)
									}
									_e.funcs.GenerateLookupJoins(_root, opt.LeftJoinOp, left, scanPrivate, _e.funcs.ConcatFilters(on, filters), private)
									if _e.o.appliedRule != nil {
										_e.o.appliedRule(opt.GenerateLookupJoinsWithFilter, _root, _last.NextExpr())
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreRightJoin(
	_rootState *exploreState,
	_root *memo.RightJoinExpr,
	_rootOrd int,
) (_fullyExplored bool) {
	_fullyExplored = true

	// [CommuteRightJoin]
	{
		if _rootOrd >= _rootState.start {
			left := _root.Left
			right := _root.Right
			on := _root.On
			private := &_root.JoinPrivate
			if _e.funcs.NoJoinHints(private) {
				if _e.o.matchedRule == nil || _e.o.matchedRule(opt.CommuteRightJoin) {
					_expr := &memo.LeftJoinExpr{
						Left:        right,
						Right:       left,
						On:          on,
						JoinPrivate: *private,
					}
					_interned := _e.mem.AddLeftJoinToGroup(_expr, _root)
					if _e.o.appliedRule != nil {
						if _interned != _expr {
							_e.o.appliedRule(opt.CommuteRightJoin, _root, nil)
						} else {
							_e.o.appliedRule(opt.CommuteRightJoin, _root, _interned)
						}
					}
				}
			}
		}
	}

	// [GenerateMergeJoins]
	{
		if _rootOrd >= _rootState.start {
			left := _root.Left
			right := _root.Right
			on := _root.On
			private := &_root.JoinPrivate
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateMergeJoins) {
				var _last memo.RelExpr
				if _e.o.appliedRule != nil {
					_last = memo.LastGroupMember(_root)
				}
				_e.funcs.GenerateMergeJoins(_root, opt.RightJoinOp, left, right, on, private)
				if _e.o.appliedRule != nil {
					_e.o.appliedRule(opt.GenerateMergeJoins, _root, _last.NextExpr())
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreFullJoin(
	_rootState *exploreState,
	_root *memo.FullJoinExpr,
	_rootOrd int,
) (_fullyExplored bool) {
	_fullyExplored = true

	// [CommuteJoin]
	{
		if _rootOrd >= _rootState.start {
			left := _root.Left
			right := _root.Right
			on := _root.On
			private := &_root.JoinPrivate
			if _e.funcs.NoJoinHints(private) {
				if _e.o.matchedRule == nil || _e.o.matchedRule(opt.CommuteJoin) {
					_expr := &memo.FullJoinExpr{
						Left:        right,
						Right:       left,
						On:          on,
						JoinPrivate: *private,
					}
					_interned := _e.mem.AddFullJoinToGroup(_expr, _root)
					if _e.o.appliedRule != nil {
						if _interned != _expr {
							_e.o.appliedRule(opt.CommuteJoin, _root, nil)
						} else {
							_e.o.appliedRule(opt.CommuteJoin, _root, _interned)
						}
					}
				}
			}
		}
	}

	// [GenerateMergeJoins]
	{
		if _rootOrd >= _rootState.start {
			left := _root.Left
			right := _root.Right
			on := _root.On
			private := &_root.JoinPrivate
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateMergeJoins) {
				var _last memo.RelExpr
				if _e.o.appliedRule != nil {
					_last = memo.LastGroupMember(_root)
				}
				_e.funcs.GenerateMergeJoins(_root, opt.FullJoinOp, left, right, on, private)
				if _e.o.appliedRule != nil {
					_e.o.appliedRule(opt.GenerateMergeJoins, _root, _last.NextExpr())
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreSemiJoin(
	_rootState *exploreState,
	_root *memo.SemiJoinExpr,
	_rootOrd int,
) (_fullyExplored bool) {
	_fullyExplored = true

	// [CommuteSemiJoin]
	{
		if _rootOrd >= _rootState.start {
			left := _root.Left
			right := _root.Right
			on := _root.On
			if _e.funcs.IsSimpleEquality(on) {
				private := &_root.JoinPrivate
				if _e.funcs.NoJoinHints(private) {
					if _e.o.matchedRule == nil || _e.o.matchedRule(opt.CommuteSemiJoin) {
						_expr := &memo.ProjectExpr{
							Input: _e.f.ConstructInnerJoin(
								left,
								_e.f.ConstructDistinctOn(
									right,
									memo.EmptyAggregationsExpr,
									_e.funcs.MakeGrouping(_e.funcs.IntersectionCols(_e.funcs.OutputCols(right), _e.funcs.FilterOuterCols(on))),
								),
								on,
								private,
							),
							Projections: memo.EmptyProjectionsExpr,
							Passthrough: _e.funcs.OutputCols(left),
						}
						_interned := _e.mem.AddProjectToGroup(_expr, _root)
						if _e.o.appliedRule != nil {
							if _interned != _expr {
								_e.o.appliedRule(opt.CommuteSemiJoin, _root, nil)
							} else {
								_e.o.appliedRule(opt.CommuteSemiJoin, _root, _interned)
							}
						}
					}
				}
			}
		}
	}

	// [GenerateMergeJoins]
	{
		if _rootOrd >= _rootState.start {
			left := _root.Left
			right := _root.Right
			on := _root.On
			private := &_root.JoinPrivate
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateMergeJoins) {
				var _last memo.RelExpr
				if _e.o.appliedRule != nil {
					_last = memo.LastGroupMember(_root)
				}
				_e.funcs.GenerateMergeJoins(_root, opt.SemiJoinOp, left, right, on, private)
				if _e.o.appliedRule != nil {
					_e.o.appliedRule(opt.GenerateMergeJoins, _root, _last.NextExpr())
				}
			}
		}
	}

	// [GenerateLookupJoins]
	{
		_partlyExplored := _rootOrd < _rootState.start
		left := _root.Left
		_state := _e.lookupExploreState(_root.Right)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = _root.Right.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			if !_partlyExplored || _ord >= _state.start {
				_scan, _ := _member.(*memo.ScanExpr)
				if _scan != nil {
					scanPrivate := &_scan.ScanPrivate
					if _e.funcs.IsCanonicalScan(scanPrivate) {
						on := _root.On
						private := &_root.JoinPrivate
						if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateLookupJoins) {
							var _last memo.RelExpr
							if _e.o.appliedRule != nil {
								_last = memo.LastGroupMember(_root)
							}
							_e.funcs.GenerateLookupJoins(_root, opt.SemiJoinOp, left, scanPrivate, on, private)
							if _e.o.appliedRule != nil {
								_e.o.appliedRule(opt.GenerateLookupJoins, _root, _last.NextExpr())
							}
						}
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreAntiJoin(
	_rootState *exploreState,
	_root *memo.AntiJoinExpr,
	_rootOrd int,
) (_fullyExplored bool) {
	_fullyExplored = true

	// [GenerateMergeJoins]
	{
		if _rootOrd >= _rootState.start {
			left := _root.Left
			right := _root.Right
			on := _root.On
			private := &_root.JoinPrivate
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateMergeJoins) {
				var _last memo.RelExpr
				if _e.o.appliedRule != nil {
					_last = memo.LastGroupMember(_root)
				}
				_e.funcs.GenerateMergeJoins(_root, opt.AntiJoinOp, left, right, on, private)
				if _e.o.appliedRule != nil {
					_e.o.appliedRule(opt.GenerateMergeJoins, _root, _last.NextExpr())
				}
			}
		}
	}

	// [GenerateLookupJoins]
	{
		_partlyExplored := _rootOrd < _rootState.start
		left := _root.Left
		_state := _e.lookupExploreState(_root.Right)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = _root.Right.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			if !_partlyExplored || _ord >= _state.start {
				_scan, _ := _member.(*memo.ScanExpr)
				if _scan != nil {
					scanPrivate := &_scan.ScanPrivate
					if _e.funcs.IsCanonicalScan(scanPrivate) {
						on := _root.On
						private := &_root.JoinPrivate
						if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateLookupJoins) {
							var _last memo.RelExpr
							if _e.o.appliedRule != nil {
								_last = memo.LastGroupMember(_root)
							}
							_e.funcs.GenerateLookupJoins(_root, opt.AntiJoinOp, left, scanPrivate, on, private)
							if _e.o.appliedRule != nil {
								_e.o.appliedRule(opt.GenerateLookupJoins, _root, _last.NextExpr())
							}
						}
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreGroupBy(
	_rootState *exploreState,
	_root *memo.GroupByExpr,
	_rootOrd int,
) (_fullyExplored bool) {
	_fullyExplored = true

	// [GenerateStreamingGroupBy]
	{
		if _rootOrd >= _rootState.start {
			input := _root.Input
			aggs := _root.Aggregations
			private := &_root.GroupingPrivate
			if _e.funcs.IsCanonicalGroupBy(private) {
				if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateStreamingGroupBy) {
					var _last memo.RelExpr
					if _e.o.appliedRule != nil {
						_last = memo.LastGroupMember(_root)
					}
					_e.funcs.GenerateStreamingGroupBy(_root, opt.GroupByOp, input, aggs, private)
					if _e.o.appliedRule != nil {
						_e.o.appliedRule(opt.GenerateStreamingGroupBy, _root, _last.NextExpr())
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreScalarGroupBy(
	_rootState *exploreState,
	_root *memo.ScalarGroupByExpr,
	_rootOrd int,
) (_fullyExplored bool) {
	_fullyExplored = true

	// [ReplaceMinWithLimit]
	{
		if _rootOrd >= _rootState.start {
			input := _root.Input
			if len(_root.Aggregations) == 1 {
				_item := &_root.Aggregations[0]
				_min, _ := _item.Agg.(*memo.MinExpr)
				if _min != nil {
					variable := _min.Input
					_variable, _ := variable.(*memo.VariableExpr)
					if _variable != nil {
						col := _variable.Col
						aggPrivate := &_item.ColPrivate
						groupingPrivate := &_root.GroupingPrivate
						if _e.funcs.IsCanonicalGroupBy(groupingPrivate) {
							if _e.o.matchedRule == nil || _e.o.matchedRule(opt.ReplaceMinWithLimit) {
								_expr := &memo.ScalarGroupByExpr{
									Input: _e.f.ConstructLimit(
										_e.f.ConstructSelect(
											input,
											memo.FiltersExpr{
												{
													Condition: _e.f.ConstructIsNot(
														variable,
														_e.f.ConstructNull(
															_e.funcs.AnyType(),
														),
													),
												},
											},
										),
										_e.f.ConstructConst(
											tree.NewDInt(1),
										),
										_e.funcs.MakeOrderingChoiceFromColumn(opt.MinOp, col),
									),
									Aggregations: memo.AggregationsExpr{
										{
											Agg: _e.f.ConstructConstAgg(
												variable,
											),
											ColPrivate: *aggPrivate,
										},
									},
									GroupingPrivate: *groupingPrivate,
								}
								_interned := _e.mem.AddScalarGroupByToGroup(_expr, _root)
								if _e.o.appliedRule != nil {
									if _interned != _expr {
										_e.o.appliedRule(opt.ReplaceMinWithLimit, _root, nil)
									} else {
										_e.o.appliedRule(opt.ReplaceMinWithLimit, _root, _interned)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// [ReplaceMaxWithLimit]
	{
		if _rootOrd >= _rootState.start {
			input := _root.Input
			if len(_root.Aggregations) == 1 {
				_item := &_root.Aggregations[0]
				_max, _ := _item.Agg.(*memo.MaxExpr)
				if _max != nil {
					variable := _max.Input
					_variable, _ := variable.(*memo.VariableExpr)
					if _variable != nil {
						col := _variable.Col
						aggPrivate := &_item.ColPrivate
						groupingPrivate := &_root.GroupingPrivate
						if _e.funcs.IsCanonicalGroupBy(groupingPrivate) {
							if _e.o.matchedRule == nil || _e.o.matchedRule(opt.ReplaceMaxWithLimit) {
								_expr := &memo.ScalarGroupByExpr{
									Input: _e.f.ConstructLimit(
										_e.f.ConstructSelect(
											input,
											memo.FiltersExpr{
												{
													Condition: _e.f.ConstructIsNot(
														variable,
														_e.f.ConstructNull(
															_e.funcs.AnyType(),
														),
													),
												},
											},
										),
										_e.f.ConstructConst(
											tree.NewDInt(1),
										),
										_e.funcs.MakeOrderingChoiceFromColumn(opt.MaxOp, col),
									),
									Aggregations: memo.AggregationsExpr{
										{
											Agg: _e.f.ConstructConstAgg(
												variable,
											),
											ColPrivate: *aggPrivate,
										},
									},
									GroupingPrivate: *groupingPrivate,
								}
								_interned := _e.mem.AddScalarGroupByToGroup(_expr, _root)
								if _e.o.appliedRule != nil {
									if _interned != _expr {
										_e.o.appliedRule(opt.ReplaceMaxWithLimit, _root, nil)
									} else {
										_e.o.appliedRule(opt.ReplaceMaxWithLimit, _root, _interned)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreDistinctOn(
	_rootState *exploreState,
	_root *memo.DistinctOnExpr,
	_rootOrd int,
) (_fullyExplored bool) {
	_fullyExplored = true

	// [GenerateStreamingGroupBy]
	{
		if _rootOrd >= _rootState.start {
			input := _root.Input
			aggs := _root.Aggregations
			private := &_root.GroupingPrivate
			if _e.funcs.IsCanonicalGroupBy(private) {
				if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateStreamingGroupBy) {
					var _last memo.RelExpr
					if _e.o.appliedRule != nil {
						_last = memo.LastGroupMember(_root)
					}
					_e.funcs.GenerateStreamingGroupBy(_root, opt.DistinctOnOp, input, aggs, private)
					if _e.o.appliedRule != nil {
						_e.o.appliedRule(opt.GenerateStreamingGroupBy, _root, _last.NextExpr())
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreLimit(
	_rootState *exploreState,
	_root *memo.LimitExpr,
	_rootOrd int,
) (_fullyExplored bool) {
	_fullyExplored = true

	// [GenerateLimitedScans]
	{
		_partlyExplored := _rootOrd < _rootState.start
		_state := _e.lookupExploreState(_root.Input)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = _root.Input.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			if !_partlyExplored || _ord >= _state.start {
				_scan, _ := _member.(*memo.ScanExpr)
				if _scan != nil {
					scanPrivate := &_scan.ScanPrivate
					if _e.funcs.IsCanonicalScan(scanPrivate) {
						_const, _ := _root.Limit.(*memo.ConstExpr)
						if _const != nil {
							limit := _const.Value
							if _e.funcs.IsPositiveLimit(limit) {
								ordering := _root.Ordering
								if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateLimitedScans) {
									var _last memo.RelExpr
									if _e.o.appliedRule != nil {
										_last = memo.LastGroupMember(_root)
									}
									_e.funcs.GenerateLimitedScans(_root, scanPrivate, limit, ordering)
									if _e.o.appliedRule != nil {
										_e.o.appliedRule(opt.GenerateLimitedScans, _root, _last.NextExpr())
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// [PushLimitIntoConstrainedScan]
	{
		_partlyExplored := _rootOrd < _rootState.start
		_state := _e.lookupExploreState(_root.Input)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = _root.Input.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			if !_partlyExplored || _ord >= _state.start {
				_scan, _ := _member.(*memo.ScanExpr)
				if _scan != nil {
					scanPrivate := &_scan.ScanPrivate
					_const, _ := _root.Limit.(*memo.ConstExpr)
					if _const != nil {
						limit := _const.Value
						if _e.funcs.IsPositiveLimit(limit) {
							ordering := _root.Ordering
							if _e.funcs.CanLimitConstrainedScan(scanPrivate, ordering) {
								if _e.o.matchedRule == nil || _e.o.matchedRule(opt.PushLimitIntoConstrainedScan) {
									_expr := &memo.ScanExpr{
										ScanPrivate: *_e.funcs.LimitScanPrivate(scanPrivate, limit, ordering),
									}
									_interned := _e.mem.AddScanToGroup(_expr, _root)
									if _e.o.appliedRule != nil {
										if _interned != _expr {
											_e.o.appliedRule(opt.PushLimitIntoConstrainedScan, _root, nil)
										} else {
											_e.o.appliedRule(opt.PushLimitIntoConstrainedScan, _root, _interned)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// [PushLimitIntoIndexJoin]
	{
		_partlyExplored := _rootOrd < _rootState.start
		_state := _e.lookupExploreState(_root.Input)
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		var _member memo.RelExpr
		for _ord := 0; _ord < _state.end; _ord++ {
			if _member == nil {
				_member = _root.Input.FirstExpr()
			} else {
				_member = _member.NextExpr()
			}
			_partlyExplored := _partlyExplored && _ord < _state.start
			_indexJoin, _ := _member.(*memo.IndexJoinExpr)
			if _indexJoin != nil {
				_state := _e.lookupExploreState(_indexJoin.Input)
				if !_state.fullyExplored {
					_fullyExplored = false
				}
				var _member memo.RelExpr
				for _ord := 0; _ord < _state.end; _ord++ {
					if _member == nil {
						_member = _indexJoin.Input.FirstExpr()
					} else {
						_member = _member.NextExpr()
					}
					if !_partlyExplored || _ord >= _state.start {
						_scan, _ := _member.(*memo.ScanExpr)
						if _scan != nil {
							scanPrivate := &_scan.ScanPrivate
							indexJoinPrivate := &_indexJoin.IndexJoinPrivate
							_const, _ := _root.Limit.(*memo.ConstExpr)
							if _const != nil {
								limit := _const.Value
								if _e.funcs.IsPositiveLimit(limit) {
									ordering := _root.Ordering
									if _e.funcs.CanLimitConstrainedScan(scanPrivate, ordering) {
										if _e.o.matchedRule == nil || _e.o.matchedRule(opt.PushLimitIntoIndexJoin) {
											_expr := &memo.IndexJoinExpr{
												Input: _e.f.ConstructScan(
													_e.funcs.LimitScanPrivate(scanPrivate, limit, ordering),
												),
												IndexJoinPrivate: *indexJoinPrivate,
											}
											_interned := _e.mem.AddIndexJoinToGroup(_expr, _root)
											if _e.o.appliedRule != nil {
												if _interned != _expr {
													_e.o.appliedRule(opt.PushLimitIntoIndexJoin, _root, nil)
												} else {
													_e.o.appliedRule(opt.PushLimitIntoIndexJoin, _root, _interned)
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return _fullyExplored
}
