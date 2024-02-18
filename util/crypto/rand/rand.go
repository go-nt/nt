package rand

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
)

// Numbers 数字字符串
func Numbers(n int) string {
	return Create(n, "0123456789")
}

// UppercaseLetters 英文大写
func UppercaseLetters(n int) string {
	return Create(n, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

// LowercaseLetters 英文小写
func LowercaseLetters(n int) string {
	return Create(n, "abcdefghijklmnopqrstuvwxyz")
}

// Letters 英文小写
func Letters(n int) string {
	return Create(n, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
}

// Simple 普通字符串，英文小写+数字
func Simple(n int) string {
	return Create(n, "abcdefghijklmnopqrstuvwxyz0123456789")
}

// Complex 复杂字符串，英文大写+英文小写+数字
func Complex(n int) string {
	return Create(n, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
}

// Secure 安全字符串，英文大写+英文小写+数字+特殊字符
func Secure(n int) string {
	return Create(n, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=")
}

// Create 用给定的字符串列表，生成随机字符串
func Create(n int, letters string) string {
	var runes []rune
	runes = []rune(letters)

	var randBytes bytes.Buffer
	randBytes.Grow(n)
	l := uint32(len(runes))
	for i := 0; i < n; i++ {
		randBytes.WriteRune(runes[binary.BigEndian.Uint32(Bytes(4))%l])
	}
	return randBytes.String()
}

// Bytes 生成一个随机字节
func Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
