// Code generated by execgen; DO NOT EDIT.
// Copyright 2019 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package exec

import (
	"context"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/execgen"
	"github.com/pkg/errors"
)

// Use execgen package to remove unused import warning.
var _ interface{} = execgen.UNSAFEGET

// NewConstOp creates a new operator that produces a constant value constVal of
// type t at index outputIdx.
func NewConstOp(
	input Operator, t coltypes.T, constVal interface{}, outputIdx int,
) (Operator, error) {
	switch t {
	case coltypes.Bool:
		return &constBoolOp{
			OneInputNode: NewOneInputNode(input),
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(bool),
		}, nil
	case coltypes.Bytes:
		return &constBytesOp{
			OneInputNode: NewOneInputNode(input),
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.([]byte),
		}, nil
	case coltypes.Decimal:
		return &constDecimalOp{
			OneInputNode: NewOneInputNode(input),
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(apd.Decimal),
		}, nil
	case coltypes.Int8:
		return &constInt8Op{
			OneInputNode: NewOneInputNode(input),
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(int8),
		}, nil
	case coltypes.Int16:
		return &constInt16Op{
			OneInputNode: NewOneInputNode(input),
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(int16),
		}, nil
	case coltypes.Int32:
		return &constInt32Op{
			OneInputNode: NewOneInputNode(input),
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(int32),
		}, nil
	case coltypes.Int64:
		return &constInt64Op{
			OneInputNode: NewOneInputNode(input),
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(int64),
		}, nil
	case coltypes.Float32:
		return &constFloat32Op{
			OneInputNode: NewOneInputNode(input),
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(float32),
		}, nil
	case coltypes.Float64:
		return &constFloat64Op{
			OneInputNode: NewOneInputNode(input),
			outputIdx:    outputIdx,
			typ:          t,
			constVal:     constVal.(float64),
		}, nil
	default:
		return nil, errors.Errorf("unsupported const type %s", t)
	}
}

type constBoolOp struct {
	OneInputNode

	typ       coltypes.T
	outputIdx int
	constVal  bool
}

func (c constBoolOp) Init() {
	c.input.Init()
}

func (c constBoolOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
	}
	if n == 0 {
		return batch
	}
	col := batch.ColVec(c.outputIdx).Bool()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel[:n] {
			col[int(i)] = c.constVal
		}
	} else {
		col = col[0:int(n)]
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constBytesOp struct {
	OneInputNode

	typ       coltypes.T
	outputIdx int
	constVal  []byte
}

func (c constBytesOp) Init() {
	c.input.Init()
}

func (c constBytesOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
	}
	if n == 0 {
		return batch
	}
	col := batch.ColVec(c.outputIdx).Bytes()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel[:n] {
			col.Set(int(i), c.constVal)
		}
	} else {
		col = col.Slice(0, int(n))
		for i := 0; i < col.Len(); i++ {
			col.Set(i, c.constVal)
		}
	}
	return batch
}

type constDecimalOp struct {
	OneInputNode

	typ       coltypes.T
	outputIdx int
	constVal  apd.Decimal
}

func (c constDecimalOp) Init() {
	c.input.Init()
}

func (c constDecimalOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
	}
	if n == 0 {
		return batch
	}
	col := batch.ColVec(c.outputIdx).Decimal()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel[:n] {
			col[int(i)].Set(&c.constVal)
		}
	} else {
		col = col[0:int(n)]
		for i := range col {
			col[i].Set(&c.constVal)
		}
	}
	return batch
}

type constInt8Op struct {
	OneInputNode

	typ       coltypes.T
	outputIdx int
	constVal  int8
}

func (c constInt8Op) Init() {
	c.input.Init()
}

func (c constInt8Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
	}
	if n == 0 {
		return batch
	}
	col := batch.ColVec(c.outputIdx).Int8()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel[:n] {
			col[int(i)] = c.constVal
		}
	} else {
		col = col[0:int(n)]
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constInt16Op struct {
	OneInputNode

	typ       coltypes.T
	outputIdx int
	constVal  int16
}

func (c constInt16Op) Init() {
	c.input.Init()
}

func (c constInt16Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
	}
	if n == 0 {
		return batch
	}
	col := batch.ColVec(c.outputIdx).Int16()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel[:n] {
			col[int(i)] = c.constVal
		}
	} else {
		col = col[0:int(n)]
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constInt32Op struct {
	OneInputNode

	typ       coltypes.T
	outputIdx int
	constVal  int32
}

func (c constInt32Op) Init() {
	c.input.Init()
}

func (c constInt32Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
	}
	if n == 0 {
		return batch
	}
	col := batch.ColVec(c.outputIdx).Int32()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel[:n] {
			col[int(i)] = c.constVal
		}
	} else {
		col = col[0:int(n)]
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constInt64Op struct {
	OneInputNode

	typ       coltypes.T
	outputIdx int
	constVal  int64
}

func (c constInt64Op) Init() {
	c.input.Init()
}

func (c constInt64Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
	}
	if n == 0 {
		return batch
	}
	col := batch.ColVec(c.outputIdx).Int64()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel[:n] {
			col[int(i)] = c.constVal
		}
	} else {
		col = col[0:int(n)]
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constFloat32Op struct {
	OneInputNode

	typ       coltypes.T
	outputIdx int
	constVal  float32
}

func (c constFloat32Op) Init() {
	c.input.Init()
}

func (c constFloat32Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
	}
	if n == 0 {
		return batch
	}
	col := batch.ColVec(c.outputIdx).Float32()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel[:n] {
			col[int(i)] = c.constVal
		}
	} else {
		col = col[0:int(n)]
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constFloat64Op struct {
	OneInputNode

	typ       coltypes.T
	outputIdx int
	constVal  float64
}

func (c constFloat64Op) Init() {
	c.input.Init()
}

func (c constFloat64Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
	}
	if n == 0 {
		return batch
	}
	col := batch.ColVec(c.outputIdx).Float64()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel[:n] {
			col[int(i)] = c.constVal
		}
	} else {
		col = col[0:int(n)]
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

// NewConstNullOp creates a new operator that produces a constant (untyped) NULL
// value at index outputIdx.
func NewConstNullOp(input Operator, outputIdx int) Operator {
	return &constNullOp{
		OneInputNode: NewOneInputNode(input),
		outputIdx:    outputIdx,
	}
}

type constNullOp struct {
	OneInputNode
	outputIdx int
}

func (c constNullOp) Init() {
	c.input.Init()
}

func (c constNullOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if batch.Width() == c.outputIdx {
		batch.AppendCol(coltypes.Int8)
	}
	if n == 0 {
		return batch
	}

	col := batch.ColVec(c.outputIdx)
	nulls := col.Nulls()
	if sel := batch.Selection(); sel != nil {
		for _, i := range sel[:n] {
			nulls.SetNull(i)
		}
	} else {
		nulls.SetNulls()
	}
	return batch
}
