// Code generated by execgen; DO NOT EDIT.

package exec

import (
	"fmt"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/types"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
)

// EncDatumRowsToColVec converts one column from EncDatumRows to a column
// vector. columnIdx is the 0-based index of the column in the EncDatumRows.
func EncDatumRowsToColVec(
	rows sqlbase.EncDatumRows,
	vec ColVec,
	columnIdx int,
	columnType *sqlbase.ColumnType,
	alloc *sqlbase.DatumAlloc,
) error {
	nRows := uint16(len(rows))
	// TODO(solon): Make this chain of conditionals more efficient: either a
	// switch statement or even better a lookup table on SemanticType. Also get
	// rid of the somewhat dubious assumption that Width is unset (0) for column
	// types where it does not apply.

	if columnType.SemanticType == sqlbase.ColumnType_NAME && columnType.Width == 0 {
		col := vec.Bytes()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.([]byte)
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_STRING && columnType.Width == 0 {
		col := vec.Bytes()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.([]byte)
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_FLOAT && columnType.Width == 0 {
		col := vec.Float64()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(float64)
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_DATE && columnType.Width == 0 {
		col := vec.Int64()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(int64)
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_OID && columnType.Width == 0 {
		col := vec.Int64()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(int64)
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_BOOL && columnType.Width == 0 {
		col := vec.Bool()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(bool)
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_INT && columnType.Width == 0 {
		col := vec.Int64()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(int64)
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_INT && columnType.Width == 8 {
		col := vec.Int8()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(int8)
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_INT && columnType.Width == 16 {
		col := vec.Int16()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(int16)
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_INT && columnType.Width == 32 {
		col := vec.Int32()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(int32)
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_INT && columnType.Width == 64 {
		col := vec.Int64()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(int64)
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_BYTES && columnType.Width == 0 {
		col := vec.Bytes()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.([]byte)
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_DECIMAL && columnType.Width == 0 {
		col := vec.Decimal()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(apd.Decimal)
			}
		}
		return nil
	}

	panic(fmt.Sprintf("Unsupported column type and width: %s, %d", columnType.SQLString(), columnType.Width))
}
