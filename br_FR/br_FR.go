package br_FR

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type br_FR struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
	percentSuffix          []byte
	perMille               []byte
	timeSeparator          []byte
	inifinity              []byte
	currencies             [][]byte // idx = enum of currency code
	currencyPositiveSuffix []byte
	currencyNegativeSuffix []byte
	monthsAbbreviated      [][]byte
	monthsNarrow           [][]byte
	monthsShort            [][]byte
	monthsWide             [][]byte
	daysAbbreviated        [][]byte
	daysNarrow             [][]byte
	daysShort              [][]byte
	daysWide               [][]byte
	periodsAbbreviated     [][]byte
	periodsNarrow          [][]byte
	periodsShort           [][]byte
	periodsWide            [][]byte
	erasAbbreviated        [][]byte
	erasNarrow             [][]byte
	erasWide               [][]byte
	timezones              map[string][]byte
}

// New returns a new instance of translator for the 'br_FR' locale
func New() locales.Translator {
	return &br_FR{
		locale:                 "br_FR",
		pluralsCardinal:        []locales.PluralRule{2, 3, 4, 5, 6},
		pluralsOrdinal:         nil,
		decimal:                []byte{},
		group:                  []byte{},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x47, 0x65, 0x6e, 0x2e}, {0x43, 0xca, 0xbc, 0x68, 0x77, 0x65, 0x2e}, {0x4d, 0x65, 0x75, 0x72, 0x2e}, {0x45, 0x62, 0x72, 0x2e}, {0x4d, 0x61, 0x65}, {0x4d, 0x65, 0x7a, 0x68, 0x2e}, {0x47, 0x6f, 0x75, 0x65, 0x2e}, {0x45, 0x6f, 0x73, 0x74}, {0x47, 0x77, 0x65, 0x6e, 0x2e}, {0x48, 0x65, 0x72, 0x65}, {0x44, 0x75}, {0x4b, 0x7a, 0x75, 0x2e}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x30, 0x31}, {0x30, 0x32}, {0x30, 0x33}, {0x30, 0x34}, {0x30, 0x35}, {0x30, 0x36}, {0x30, 0x37}, {0x30, 0x38}, {0x30, 0x39}, {0x31, 0x30}, {0x31, 0x31}, {0x31, 0x32}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x47, 0x65, 0x6e, 0x76, 0x65, 0x72}, {0x43, 0xca, 0xbc, 0x68, 0x77, 0x65, 0x76, 0x72, 0x65, 0x72}, {0x4d, 0x65, 0x75, 0x72, 0x7a, 0x68}, {0x45, 0x62, 0x72, 0x65, 0x6c}, {0x4d, 0x61, 0x65}, {0x4d, 0x65, 0x7a, 0x68, 0x65, 0x76, 0x65, 0x6e}, {0x47, 0x6f, 0x75, 0x65, 0x72, 0x65}, {0x45, 0x6f, 0x73, 0x74}, {0x47, 0x77, 0x65, 0x6e, 0x67, 0x6f, 0x6c, 0x6f}, {0x48, 0x65, 0x72, 0x65}, {0x44, 0x75}, {0x4b, 0x65, 0x72, 0x7a, 0x75}},
		daysAbbreviated:        [][]uint8{{0x53, 0x75, 0x6c}, {0x4c, 0x75, 0x6e}, {0x4d, 0x65, 0x75, 0x2e}, {0x4d, 0x65, 0x72, 0x2e}, {0x59, 0x61, 0x6f, 0x75}, {0x47, 0x77, 0x65, 0x2e}, {0x53, 0x61, 0x64, 0x2e}},
		daysNarrow:             [][]uint8{{0x53, 0x75}, {0x4c}, {0x4d, 0x7a}, {0x4d, 0x63}, {0x59}, {0x47}, {0x53, 0x61}},
		daysWide:               [][]uint8{{0x53, 0x75, 0x6c}, {0x4c, 0x75, 0x6e}, {0x4d, 0x65, 0x75, 0x72, 0x7a, 0x68}, {0x4d, 0x65, 0x72, 0x63, 0xca, 0xbc, 0x68, 0x65, 0x72}, {0x59, 0x61, 0x6f, 0x75}, {0x47, 0x77, 0x65, 0x6e, 0x65, 0x72}, {0x53, 0x61, 0x64, 0x6f, 0x72, 0x6e}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x2e, 0x4d, 0x2e}, {0x47, 0x2e, 0x4d, 0x2e}},
		periodsNarrow:          [][]uint8{{0x61, 0x6d}, {0x67, 0x6d}},
		periodsWide:            [][]uint8{{0x41, 0x2e, 0x4d, 0x2e}, {0x47, 0x2e, 0x4d, 0x2e}},
		erasAbbreviated:        [][]uint8{{0x61, 0x2d, 0x72, 0x61, 0x6f, 0x6b, 0x20, 0x4a, 0x2e, 0x4b, 0x2e}, {0x67, 0x6f, 0x75, 0x64, 0x65, 0x20, 0x4a, 0x2e, 0x4b, 0x2e}},
		erasNarrow:             [][]uint8{{0x61, 0x2d, 0x72, 0x61, 0x6f, 0x6b, 0x20, 0x4a, 0x2e, 0x4b, 0x2e}, {0x67, 0x6f, 0x75, 0x64, 0x65, 0x20, 0x4a, 0x2e, 0x4b, 0x2e}},
		erasWide:               [][]uint8{{0x61, 0x2d, 0x72, 0x61, 0x6f, 0x6b, 0x20, 0x4a, 0x65, 0x7a, 0x75, 0x7a, 0x2d, 0x4b, 0x72, 0x69, 0x73, 0x74}, {0x67, 0x6f, 0x75, 0x64, 0x65, 0x20, 0x4a, 0x65, 0x7a, 0x75, 0x7a, 0x2d, 0x4b, 0x72, 0x69, 0x73, 0x74}},
		timezones:              map[string][]uint8{"PST": {0x50, 0x53, 0x54}, "GMT": {0x41, 0x6d, 0x7a, 0x65, 0x72, 0x20, 0x6b, 0x65, 0x69, 0x74, 0x61, 0x74, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x20, 0x28, 0x41, 0x4b, 0x47, 0x29}, "AKST": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "ACST": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4b, 0x72, 0x65, 0x69, 0x7a, 0x61, 0x6f, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61}, "JDT": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4a, 0x61, 0x70, 0x61, 0x6e}, "CHADT": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "ART": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x41, 0x72, 0x63, 0xca, 0xbc, 0x68, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "CAT": {0x65, 0x75, 0x72, 0x20, 0x4b, 0x72, 0x65, 0x69, 0x7a, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61}, "WIT": {0x65, 0x75, 0x72, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x61, 0x20, 0x61, 0x72, 0x20, 0x52, 0x65, 0x74, 0x65, 0x72}, "AST": {0x41, 0x53, 0x54}, "MEZ": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4b, 0x72, 0x65, 0x69, 0x7a, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x61}, "MESZ": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4b, 0x72, 0x65, 0x69, 0x7a, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x61}, "BT": {0x65, 0x75, 0x72, 0x20, 0x42, 0x68, 0x6f, 0x75, 0x74, 0x61, 0x6e}, "MST": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4d, 0x61, 0x63, 0x61, 0x75}, "COT": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4b, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "SAST": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x53, 0x75, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61}, "IST": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x49, 0x6e, 0x64, 0x69, 0x61}, "TMST": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "ARST": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x41, 0x72, 0x63, 0xca, 0xbc, 0x68, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "WART": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x41, 0x72, 0x63, 0xca, 0xbc, 0x68, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x61, 0x72, 0x20, 0x43, 0xca, 0xbc, 0x68, 0x6f, 0x72, 0x6e, 0xc3, 0xb4, 0x67}, "ACDT": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4b, 0x72, 0x65, 0x69, 0x7a, 0x61, 0x6f, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61}, "EST": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x61, 0x72, 0x20, 0x52, 0x65, 0x74, 0x65, 0x72}, "HKST": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "CDT": {0x43, 0x44, 0x54}, "NZST": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x2d, 0x4e, 0x65, 0x76, 0x65, 0x7a}, "TMT": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "HAT": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "PDT": {0x50, 0x44, 0x54}, "MYT": {0x65, 0x75, 0x72, 0x20, 0x4d, 0x61, 0x6c, 0x61, 0x79, 0x73, 0x69, 0x61}, "ADT": {0x41, 0x44, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "JST": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4a, 0x61, 0x70, 0x61, 0x6e}, "UYST": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "ChST": {0x43, 0x68, 0x53, 0x54}, "AEDT": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x41, 0x6f, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x61, 0x72, 0x20, 0x52, 0x65, 0x74, 0x65, 0x72}, "NZDT": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x2d, 0x4e, 0x65, 0x76, 0x65, 0x7a}, "CLT": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "AWDT": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x41, 0x6f, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x61, 0x72, 0x20, 0x43, 0xca, 0xbc, 0x68, 0x6f, 0x72, 0x6e, 0xc3, 0xb4, 0x67}, "CHAST": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "WIB": {0x65, 0x75, 0x72, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x61, 0x20, 0x61, 0x72, 0x20, 0x43, 0xca, 0xbc, 0x68, 0x6f, 0x72, 0x6e, 0xc3, 0xb4, 0x67}, "SGT": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x6f, 0x75, 0x72}, "EAT": {0x65, 0x75, 0x72, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x61, 0x72, 0x20, 0x52, 0x65, 0x74, 0x65, 0x72}, "EDT": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x61, 0x72, 0x20, 0x52, 0x65, 0x74, 0x65, 0x72}, "AWST": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x41, 0x6f, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x61, 0x72, 0x20, 0x43, 0xca, 0xbc, 0x68, 0x6f, 0x72, 0x6e, 0xc3, 0xb4, 0x67}, "OEZ": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x61, 0x72, 0x20, 0x52, 0x65, 0x74, 0x65, 0x72}, "AKDT": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "HKT": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "BOT": {0x65, 0x75, 0x72, 0x20, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61}, "CST": {0x43, 0x53, 0x54}, "VET": {0x65, 0x75, 0x72, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61}, "HADT": {0x48, 0x41, 0x44, 0x54}, "OESZ": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x61, 0x72, 0x20, 0x52, 0x65, 0x74, 0x65, 0x72}, "COST": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4b, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "WEZ": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x61, 0x72, 0x20, 0x43, 0xca, 0xbc, 0x68, 0x6f, 0x72, 0x6e, 0xc3, 0xb4, 0x67}, "WESZ": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x61, 0x72, 0x20, 0x43, 0xca, 0xbc, 0x68, 0x6f, 0x72, 0x6e, 0xc3, 0xb4, 0x67}, "WARST": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x41, 0x72, 0x63, 0xca, 0xbc, 0x68, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x61, 0x72, 0x20, 0x43, 0xca, 0xbc, 0x68, 0x6f, 0x72, 0x6e, 0xc3, 0xb4, 0x67}, "GYT": {0x65, 0x75, 0x72, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61}, "AEST": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x41, 0x6f, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x61, 0x72, 0x20, 0x52, 0x65, 0x74, 0x65, 0x72}, "WITA": {0x57, 0x49, 0x54, 0x41}, "MDT": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4d, 0x61, 0x63, 0x61, 0x75}, "∅∅∅": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x42, 0x72, 0x61, 0x73, 0xc3, 0xad, 0x6c, 0x69, 0x61}, "UYT": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "ACWST": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4b, 0x72, 0x65, 0x69, 0x7a, 0x61, 0x6f, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x61, 0x72, 0x20, 0x43, 0xca, 0xbc, 0x68, 0x6f, 0x72, 0x6e, 0xc3, 0xb4, 0x67}, "CLST": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "HNT": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "WAT": {0x65, 0x75, 0x72, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x6f, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x61, 0x72, 0x20, 0x43, 0xca, 0xbc, 0x68, 0x6f, 0x72, 0x6e, 0xc3, 0xb4, 0x67}, "WAST": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x61, 0x72, 0x20, 0x43, 0xca, 0xbc, 0x68, 0x6f, 0x72, 0x6e, 0xc3, 0xb4, 0x67}, "GFT": {0x65, 0x75, 0x72, 0x20, 0x47, 0x77, 0x69, 0x61, 0x6e, 0x61, 0x20, 0x63, 0xca, 0xbc, 0x68, 0x61, 0x6c, 0x6c}, "SRT": {0x65, 0x75, 0x72, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d}, "ECT": {0x65, 0x75, 0x72, 0x20, 0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72}, "ACWDT": {0x65, 0x75, 0x72, 0x20, 0x68, 0x61, 0xc3, 0xb1, 0x76, 0x20, 0x4b, 0x72, 0x65, 0x69, 0x7a, 0x61, 0x6f, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x61, 0x72, 0x20, 0x43, 0xca, 0xbc, 0x68, 0x6f, 0x72, 0x6e, 0xc3, 0xb4, 0x67}},
	}
}

// Locale returns the current translators string locale
func (br *br_FR) Locale() string {
	return br.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'br_FR'
func (br *br_FR) PluralsCardinal() []locales.PluralRule {
	return br.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'br_FR'
func (br *br_FR) PluralsOrdinal() []locales.PluralRule {
	return br.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'br_FR'
func (br *br_FR) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)
	nMod1000000 := math.Mod(n, 1000000)

	if nMod10 == 1 && (nMod100 != 11 && nMod100 != 71 && nMod100 != 91) {
		return locales.PluralRuleOne
	} else if nMod10 == 2 && (nMod100 != 12 && nMod100 != 72 && nMod100 != 92) {
		return locales.PluralRuleTwo
	} else if nMod10 >= 3 && nMod10 <= 4 && (nMod10 == 9) && (nMod100 < 10 && nMod100 > 19) || (nMod100 < 70 && nMod100 > 79) || (nMod100 < 90 && nMod100 > 99) {
		return locales.PluralRuleFew
	} else if n != 0 && nMod1000000 == 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'br_FR'
func (br *br_FR) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'br_FR'
func (br *br_FR) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'br_FR' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (br *br_FR) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(br.decimal) + len(br.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(br.decimal) - 1; j >= 0; j-- {
				b = append(b, br.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(br.group) - 1; j >= 0; j-- {
					b = append(b, br.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(br.minus) - 1; j >= 0; j-- {
			b = append(b, br.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'br_FR' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (br *br_FR) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(br.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(br.decimal) - 1; j >= 0; j-- {
				b = append(b, br.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(br.minus) - 1; j >= 0; j-- {
			b = append(b, br.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, br.percentSuffix...)

	b = append(b, br.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'br_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (br *br_FR) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := br.currencies[currency]
	l := len(s) + len(br.decimal) + len(br.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(br.decimal) - 1; j >= 0; j-- {
				b = append(b, br.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(br.group) - 1; j >= 0; j-- {
					b = append(b, br.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(br.minus) - 1; j >= 0; j-- {
			b = append(b, br.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, br.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, br.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'br_FR'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (br *br_FR) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := br.currencies[currency]
	l := len(s) + len(br.decimal) + len(br.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(br.decimal) - 1; j >= 0; j-- {
				b = append(b, br.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(br.group) - 1; j >= 0; j-- {
					b = append(b, br.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(br.minus) - 1; j >= 0; j-- {
			b = append(b, br.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, br.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, br.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, br.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'br_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (br *br_FR) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'br_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (br *br_FR) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'br_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (br *br_FR) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'br_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (br *br_FR) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'br_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (br *br_FR) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'br_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (br *br_FR) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'br_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (br *br_FR) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'br_FR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (br *br_FR) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}
