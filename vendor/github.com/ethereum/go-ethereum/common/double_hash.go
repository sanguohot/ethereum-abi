// Copyright 2018 Sanguohot
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package common

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Lengths of doubleHashes and addresses in bytes.
const (
	// DoubleHashLength is the expected length of the doubleHash
	DoubleHashLength = 64
)

var (
	doubleHashT    = reflect.TypeOf(DoubleHash{})
)

// DoubleHash represents the 32 byte Keccak256 doubleHash of arbitrary data.
type DoubleHash [DoubleHashLength]byte

// BytesToDoubleHash sets b to doubleHash.
// If b is larger than len(h), b will be cropped from the left.
func BytesToDoubleHash(b []byte) DoubleHash {
	var h DoubleHash
	h.SetBytes(b)
	return h
}

// BigToDoubleHash sets byte representation of b to doubleHash.
// If b is larger than len(h), b will be cropped from the left.
func BigToDoubleHash(b *big.Int) DoubleHash { return BytesToDoubleHash(b.Bytes()) }

// HexToDoubleHash sets byte representation of s to doubleHash.
// If b is larger than len(h), b will be cropped from the left.
func HexToDoubleHash(s string) DoubleHash { return BytesToDoubleHash(FromHex(s)) }

// Bytes gets the byte representation of the underlying doubleHash.
func (h DoubleHash) Bytes() []byte { return h[:] }

// Big converts a doubleHash to a big integer.
func (h DoubleHash) Big() *big.Int { return new(big.Int).SetBytes(h[:]) }

// Hex converts a doubleHash to a hex string.
func (h DoubleHash) Hex() string { return hexutil.Encode(h[:]) }

// TerminalString implements log.TerminalStringer, formatting a string for console
// output during logging.
func (h DoubleHash) TerminalString() string {
	return fmt.Sprintf("%x…%x", h[:3], h[29:])
}

// String implements the stringer interface and is used also by the logger when
// doing full logging into a file.
func (h DoubleHash) String() string {
	return h.Hex()
}

// Format implements fmt.Formatter, forcing the byte slice to be formatted as is,
// without going through the stringer interface used for logging.
func (h DoubleHash) Format(s fmt.State, c rune) {
	fmt.Fprintf(s, "%"+string(c), h[:])
}

// UnmarshalText parses a doubleHash in hex syntax.
func (h *DoubleHash) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("DoubleHash", input, h[:])
}

// UnmarshalJSON parses a doubleHash in hex syntax.
func (h *DoubleHash) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(doubleHashT, input, h[:])
}

// MarshalText returns the hex representation of h.
func (h DoubleHash) MarshalText() ([]byte, error) {
	return hexutil.Bytes(h[:]).MarshalText()
}

// SetBytes sets the doubleHash to the value of b.
// If b is larger than len(h), b will be cropped from the left.
func (h *DoubleHash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-DoubleHashLength:]
	}

	copy(h[DoubleHashLength-len(b):], b)
}

// Generate implements testing/quick.Generator.
func (h DoubleHash) Generate(rand *rand.Rand, size int) reflect.Value {
	m := rand.Intn(len(h))
	for i := len(h) - 1; i > m; i-- {
		h[i] = byte(rand.Uint32())
	}
	return reflect.ValueOf(h)
}

// Scan implements Scanner for database/sql.
func (h *DoubleHash) Scan(src interface{}) error {
	srcB, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can't scan %T into DoubleHash", src)
	}
	if len(srcB) != DoubleHashLength {
		return fmt.Errorf("can't scan []byte of len %d into DoubleHash, want %d", len(srcB), DoubleHashLength)
	}
	copy(h[:], srcB)
	return nil
}

// Value implements valuer for database/sql.
func (h DoubleHash) Value() (driver.Value, error) {
	return h[:], nil
}

// UnprefixedDoubleHash allows marshaling a DoubleHash without 0x prefix.
type UnprefixedDoubleHash DoubleHash

// UnmarshalText decodes the doubleHash from hex. The 0x prefix is optional.
func (h *UnprefixedDoubleHash) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedUnprefixedText("UnprefixedDoubleHash", input, h[:])
}

// MarshalText encodes the doubleHash as hex.
func (h UnprefixedDoubleHash) MarshalText() ([]byte, error) {
	return []byte(hex.EncodeToString(h[:])), nil
}