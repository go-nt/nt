package request

import (
	"strconv"
)

type Format struct {
	Value string
}

// Byte 格式化为 byte
func (f *Format) Byte(defaultValue byte) byte {
	if len(f.Value) != 1 {
		return defaultValue
	}

	return f.Value[0]
}

// Int 格式化为 int
func (f *Format) Int(defaultValue int) int {
	if f.Value == "" {
		return defaultValue
	}

	intVal, err := strconv.Atoi(f.Value)
	if err != nil {
		return defaultValue
	}

	return intVal
}

// Rune 格式化为 Rune
func (f *Format) Rune(defaultValue rune) rune {
	if f.Value == "" {
		return defaultValue
	}

	intVal, err := strconv.Atoi(f.Value)
	if err != nil {
		return defaultValue
	}

	return rune(intVal)
}

// Int8 格式化为 int8
func (f *Format) Int8(defaultValue int8) int8 {
	if f.Value == "" {
		return defaultValue
	}

	val, err := strconv.ParseInt(f.Value, 10, 8)
	if err != nil {
		return defaultValue
	}

	return int8(val)
}

// Int16 格式化为 int16
func (f *Format) Int16(defaultValue int16) int16 {
	if f.Value == "" {
		return defaultValue
	}

	val, err := strconv.ParseInt(f.Value, 10, 16)
	if err != nil {
		return defaultValue
	}

	return int16(val)
}

// Int32 格式化为 int32
func (f *Format) Int32(defaultValue int32) int32 {
	if f.Value == "" {
		return defaultValue
	}

	val, err := strconv.ParseInt(f.Value, 10, 32)
	if err != nil {
		return defaultValue
	}

	return int32(val)
}

// Int64 格式化为 int64
func (f *Format) Int64(defaultValue int64) int64 {
	if f.Value == "" {
		return defaultValue
	}

	val, err := strconv.ParseInt(f.Value, 10, 64)
	if err != nil {
		return defaultValue
	}

	return val
}

// Uint 格式化为 uint
func (f *Format) Uint(defaultValue uint) uint {
	if f.Value == "" {
		return defaultValue
	}

	val, err := strconv.ParseUint(f.Value, 10, 64)
	if err != nil {
		return defaultValue
	}

	return uint(val)
}

// Uint8 格式化为 uint8
func (f *Format) Uint8(defaultValue uint8) uint8 {
	if f.Value == "" {
		return defaultValue
	}

	val, err := strconv.ParseUint(f.Value, 10, 8)
	if err != nil {
		return defaultValue
	}

	return uint8(val)
}

// Uint16 格式化为 uint16
func (f *Format) Uint16(defaultValue uint16) uint16 {
	if f.Value == "" {
		return defaultValue
	}

	val, err := strconv.ParseUint(f.Value, 10, 16)
	if err != nil {
		return defaultValue
	}

	return uint16(val)
}

// Uint32 格式化为 uint32
func (f *Format) Uint32(defaultValue uint32) uint32 {
	if f.Value == "" {
		return defaultValue
	}

	val, err := strconv.ParseUint(f.Value, 10, 32)
	if err != nil {
		return defaultValue
	}

	return uint32(val)
}

// Unt64 格式化为 uint64
func (f *Format) Unt64(defaultValue uint64) uint64 {
	if f.Value == "" {
		return defaultValue
	}

	val, err := strconv.ParseUint(f.Value, 10, 64)
	if err != nil {
		return defaultValue
	}

	return val
}

// Float32 格式化为 float32
func (f *Format) Float32(defaultValue float32) float32 {
	if f.Value == "" {
		return defaultValue
	}

	val, err := strconv.ParseFloat(f.Value, 32)
	if err != nil {
		return defaultValue
	}

	return float32(val)
}

// Float64 格式化为 float64
func (f *Format) Float64(defaultValue float64) float64 {
	if f.Value == "" {
		return defaultValue
	}

	val, err := strconv.ParseFloat(f.Value, 64)
	if err != nil {
		return defaultValue
	}

	return val
}
