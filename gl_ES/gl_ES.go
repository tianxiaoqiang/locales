package gl_ES

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type gl_ES struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
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

// New returns a new instance of translator for the 'gl_ES' locale
func New() locales.Translator {
	return &gl_ES{
		locale:                 "gl_ES",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x78, 0x61, 0x6e}, {0x66, 0x65, 0x62}, {0x6d, 0x61, 0x72}, {0x61, 0x62, 0x72}, {0x6d, 0x61, 0x69}, {0x78, 0x75, 0xc3, 0xb1}, {0x78, 0x75, 0x6c}, {0x61, 0x67, 0x6f}, {0x73, 0x65, 0x74}, {0x6f, 0x75, 0x74}, {0x6e, 0x6f, 0x76}, {0x64, 0x65, 0x63}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x58}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x58}, {0x58}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x78, 0x61, 0x6e, 0x65, 0x69, 0x72, 0x6f}, {0x66, 0x65, 0x62, 0x72, 0x65, 0x69, 0x72, 0x6f}, {0x6d, 0x61, 0x72, 0x7a, 0x6f}, {0x61, 0x62, 0x72, 0x69, 0x6c}, {0x6d, 0x61, 0x69, 0x6f}, {0x78, 0x75, 0xc3, 0xb1, 0x6f}, {0x78, 0x75, 0x6c, 0x6c, 0x6f}, {0x61, 0x67, 0x6f, 0x73, 0x74, 0x6f}, {0x73, 0x65, 0x74, 0x65, 0x6d, 0x62, 0x72, 0x6f}, {0x6f, 0x75, 0x74, 0x75, 0x62, 0x72, 0x6f}, {0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x72, 0x6f}, {0x64, 0x65, 0x63, 0x65, 0x6d, 0x62, 0x72, 0x6f}},
		daysAbbreviated:        [][]uint8{{0x64, 0x6f, 0x6d}, {0x6c, 0x75, 0x6e, 0x73}, {0x6d, 0x61, 0x72}, {0x6d, 0xc3, 0xa9, 0x72}, {0x78, 0x6f, 0x76}, {0x76, 0x65, 0x6e}, {0x73, 0xc3, 0xa1, 0x62}},
		daysNarrow:             [][]uint8{{0x44}, {0x4c}, {0x4d}, {0x4d}, {0x58}, {0x56}, {0x53}},
		daysShort:              [][]uint8{{0x64, 0x6f, 0x6d}, {0x6c, 0x75, 0x6e, 0x73}, {0x6d, 0x74}, {0x6d, 0x63}, {0x78, 0x76}, {0x76, 0x65}, {0x73, 0xc3, 0xa1, 0x62}},
		daysWide:               [][]uint8{{0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x6f}, {0x6c, 0x75, 0x6e, 0x73}, {0x6d, 0x61, 0x72, 0x74, 0x65, 0x73}, {0x6d, 0xc3, 0xa9, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x73}, {0x78, 0x6f, 0x76, 0x65, 0x73}, {0x76, 0x65, 0x6e, 0x72, 0x65, 0x73}, {0x73, 0xc3, 0xa1, 0x62, 0x61, 0x64, 0x6f}},
		periodsAbbreviated:     [][]uint8{{0x61, 0x2e, 0x6d, 0x2e}, {0x70, 0x2e, 0x6d, 0x2e}},
		periodsNarrow:          [][]uint8{{0x61, 0x2e, 0x6d, 0x2e}, {0x70, 0x2e, 0x6d, 0x2e}},
		periodsWide:            [][]uint8{{0x61, 0x2e, 0x6d, 0x2e}, {0x70, 0x2e, 0x6d, 0x2e}},
		erasAbbreviated:        [][]uint8{{0x61, 0x2e, 0x43, 0x2e}, {0x64, 0x2e, 0x43, 0x2e}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x61, 0x6e, 0x74, 0x65, 0x73, 0x20, 0x64, 0x65, 0x20, 0x43, 0x72, 0x69, 0x73, 0x74, 0x6f}, {0x64, 0x65, 0x73, 0x70, 0x6f, 0x69, 0x73, 0x20, 0x64, 0x65, 0x20, 0x43, 0x72, 0x69, 0x73, 0x74, 0x6f}},
		timezones:              map[string][]uint8{"EDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x6d, 0xc3, 0xa9, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "AST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x6f, 0x20, 0x41, 0x74, 0x6c, 0xc3, 0xa1, 0x6e, 0x74, 0x69, 0x63, 0x6f}, "HADT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x6f}, "LHDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "GFT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x47, 0xc3, 0xbc, 0x69, 0x61, 0x6e, 0x61, 0x20, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x61}, "WEZ": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "AWDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "JDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x58, 0x61, 0x70, 0xc3, 0xb3, 0x6e}, "SGT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72}, "WESZ": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "HKST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "SAST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x53, 0x75, 0x64, 0xc3, 0xa1, 0x66, 0x72, 0x69, 0x63, 0x61}, "WAT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "COT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "NZST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x6f, 0x76, 0x61, 0x20, 0x43, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x61}, "ADT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x6f, 0x20, 0x41, 0x74, 0x6c, 0xc3, 0xa1, 0x6e, 0x74, 0x69, 0x63, 0x6f}, "WIB": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "MEZ": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "CLST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "VET": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61}, "WIT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "TMT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x54, 0x75, 0x72, 0x63, 0x6f, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0xc3, 0xa1, 0x6e}, "UYST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69}, "AKST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x63, 0x61}, "ACDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "WITA": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "JST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x58, 0x61, 0x70, 0xc3, 0xb3, 0x6e}, "HKT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "COST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "SRT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d}, "LHST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "OESZ": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "BOT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61}, "ACWDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "ACST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "TMST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x54, 0x75, 0x72, 0x63, 0x6f, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0xc3, 0xa1, 0x6e}, "HNT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x6f, 0x76, 0x61}, "EAT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "CLT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "NZDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x6f, 0x76, 0x61, 0x20, 0x43, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x61}, "ACWST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "IST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x61, 0x20, 0x49, 0x6e, 0x64, 0x69, 0x61}, "OEZ": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "CHADT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "GMT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x69, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68}, "UYT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69}, "WART": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x72, 0x78, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "WARST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x72, 0x78, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "AEDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "MST": {0x4d, 0x53, 0x54}, "MESZ": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "AWST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "BT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x42, 0x75, 0x74, 0xc3, 0xa1, 0x6e}, "CDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x61, 0x20, 0x7a, 0x6f, 0x6e, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "MDT": {0x4d, 0x44, 0x54}, "PDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x6f, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63, 0x6f}, "ARST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x72, 0x78, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "CST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "EST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x41, 0x6d, 0xc3, 0xa9, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "GYT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x47, 0xc3, 0xbc, 0x69, 0x61, 0x6e, 0x61}, "WAST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "ART": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x72, 0x78, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "ECT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72}, "AEST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "HAST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x6f}, "∅∅∅": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x42, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x69, 0x61}, "ChST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "HAT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x6f, 0x76, 0x61}, "CAT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "CHAST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "PST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x6f, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63, 0x6f}, "MYT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x4d, 0x61, 0x6c, 0x61, 0x69, 0x73, 0x69, 0x61}, "AKDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x63, 0x61}},
	}
}

// Locale returns the current translators string locale
func (gl *gl_ES) Locale() string {
	return gl.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'gl_ES'
func (gl *gl_ES) PluralsCardinal() []locales.PluralRule {
	return gl.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'gl_ES'
func (gl *gl_ES) PluralsOrdinal() []locales.PluralRule {
	return gl.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'gl_ES'
func (gl *gl_ES) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'gl_ES'
func (gl *gl_ES) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'gl_ES'
func (gl *gl_ES) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := gl.CardinalPluralRule(num1, v1)
	end := gl.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'gl_ES' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl_ES) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(gl.decimal) + len(gl.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, gl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, gl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'gl_ES' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (gl *gl_ES) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(gl.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gl.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, gl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, gl.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'gl_ES'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl_ES) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := gl.currencies[currency]
	l := len(s) + len(gl.decimal) + len(gl.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, gl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, gl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, gl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, gl.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'gl_ES'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl_ES) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := gl.currencies[currency]
	l := len(s) + len(gl.decimal) + len(gl.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, gl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, gl.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, gl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, gl.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, gl.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'gl_ES'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl_ES) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'gl_ES'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl_ES) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, gl.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'gl_ES'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl_ES) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, gl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'gl_ES'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl_ES) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, gl.daysWide[t.Day()]...)
	b = append(b, []byte{0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, gl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'gl_ES'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl_ES) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'gl_ES'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl_ES) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'gl_ES'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl_ES) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'gl_ES'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl_ES) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := gl.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
