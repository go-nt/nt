package session

type Format struct {
	Value any
}

// String 格式化为 string
func (f *Format) String(defaultValue string) string {
	if val, ok := f.Value.(string); ok {
		return val
	}

	return defaultValue
}

// Byte 格式化为 byte
func (f *Format) Byte(defaultValue byte) byte {
	if val, ok := f.Value.(byte); ok {
		return val
	}

	return defaultValue
}

// Int 格式化为 int
func (f *Format) Int(defaultValue int) int {
	if val, ok := f.Value.(int); ok {
		return val
	}

	return defaultValue
}

// Rune 格式化为 Rune
func (f *Format) Rune(defaultValue rune) rune {
	if val, ok := f.Value.(rune); ok {
		return val
	}

	return defaultValue
}

// Int8 格式化为 int8
func (f *Format) Int8(defaultValue int8) int8 {
	if val, ok := f.Value.(int8); ok {
		return val
	}

	return defaultValue
}

// Int16 格式化为 int16
func (f *Format) Int16(defaultValue int16) int16 {
	if val, ok := f.Value.(int16); ok {
		return val
	}

	return defaultValue
}

// Int32 格式化为 int32
func (f *Format) Int32(defaultValue int32) int32 {
	if val, ok := f.Value.(int32); ok {
		return val
	}

	return defaultValue
}

// Int64 格式化为 int64
func (f *Format) Int64(defaultValue int64) int64 {
	if val, ok := f.Value.(int64); ok {
		return val
	}

	return defaultValue
}

// Uint 格式化为 uint
func (f *Format) Uint(defaultValue uint) uint {
	if val, ok := f.Value.(uint); ok {
		return val
	}

	return defaultValue
}

// Uint8 格式化为 uint8
func (f *Format) Uint8(defaultValue uint8) uint8 {
	if val, ok := f.Value.(uint8); ok {
		return val
	}

	return defaultValue
}

// Uint16 格式化为 uint16
func (f *Format) Uint16(defaultValue uint16) uint16 {
	if val, ok := f.Value.(uint16); ok {
		return val
	}

	return defaultValue
}

// Uint32 格式化为 uint32
func (f *Format) Uint32(defaultValue uint32) uint32 {
	if val, ok := f.Value.(uint32); ok {
		return val
	}

	return defaultValue
}

// Unt64 格式化为 uint64
func (f *Format) Unt64(defaultValue uint64) uint64 {
	if val, ok := f.Value.(uint64); ok {
		return val
	}

	return defaultValue
}

// Float32 格式化为 float32
func (f *Format) Float32(defaultValue float32) float32 {
	if val, ok := f.Value.(float32); ok {
		return val
	}

	return defaultValue
}

// Float64 格式化为 float64
func (f *Format) Float64(defaultValue float64) float64 {
	if val, ok := f.Value.(float64); ok {
		return val
	}

	return defaultValue
}
