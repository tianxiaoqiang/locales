package hr_HR

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type hr_HR struct {
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

// New returns a new instance of translator for the 'hr_HR' locale
func New() locales.Translator {
	return &hr_HR{
		locale:                 "hr_HR",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
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
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x73, 0x69, 0x6a}, {0x76, 0x65, 0x6c, 0x6a}, {0x6f, 0xc5, 0xbe, 0x75}, {0x74, 0x72, 0x61}, {0x73, 0x76, 0x69}, {0x6c, 0x69, 0x70}, {0x73, 0x72, 0x70}, {0x6b, 0x6f, 0x6c}, {0x72, 0x75, 0x6a}, {0x6c, 0x69, 0x73}, {0x73, 0x74, 0x75}, {0x70, 0x72, 0x6f}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x31, 0x2e}, {0x32, 0x2e}, {0x33, 0x2e}, {0x34, 0x2e}, {0x35, 0x2e}, {0x36, 0x2e}, {0x37, 0x2e}, {0x38, 0x2e}, {0x39, 0x2e}, {0x31, 0x30, 0x2e}, {0x31, 0x31, 0x2e}, {0x31, 0x32, 0x2e}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x73, 0x69, 0x6a, 0x65, 0xc4, 0x8d, 0x6e, 0x6a, 0x61}, {0x76, 0x65, 0x6c, 0x6a, 0x61, 0xc4, 0x8d, 0x65}, {0x6f, 0xc5, 0xbe, 0x75, 0x6a, 0x6b, 0x61}, {0x74, 0x72, 0x61, 0x76, 0x6e, 0x6a, 0x61}, {0x73, 0x76, 0x69, 0x62, 0x6e, 0x6a, 0x61}, {0x6c, 0x69, 0x70, 0x6e, 0x6a, 0x61}, {0x73, 0x72, 0x70, 0x6e, 0x6a, 0x61}, {0x6b, 0x6f, 0x6c, 0x6f, 0x76, 0x6f, 0x7a, 0x61}, {0x72, 0x75, 0x6a, 0x6e, 0x61}, {0x6c, 0x69, 0x73, 0x74, 0x6f, 0x70, 0x61, 0x64, 0x61}, {0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x6f, 0x67, 0x61}, {0x70, 0x72, 0x6f, 0x73, 0x69, 0x6e, 0x63, 0x61}},
		daysAbbreviated:        [][]uint8{{0x6e, 0x65, 0x64}, {0x70, 0x6f, 0x6e}, {0x75, 0x74, 0x6f}, {0x73, 0x72, 0x69}, {0xc4, 0x8d, 0x65, 0x74}, {0x70, 0x65, 0x74}, {0x73, 0x75, 0x62}},
		daysNarrow:             [][]uint8{{0x4e}, {0x50}, {0x55}, {0x53}, {0xc4, 0x8c}, {0x50}, {0x53}},
		daysShort:              [][]uint8{{0x6e, 0x65, 0x64}, {0x70, 0x6f, 0x6e}, {0x75, 0x74, 0x6f}, {0x73, 0x72, 0x69}, {0xc4, 0x8d, 0x65, 0x74}, {0x70, 0x65, 0x74}, {0x73, 0x75, 0x62}},
		daysWide:               [][]uint8{{0x6e, 0x65, 0x64, 0x6a, 0x65, 0x6c, 0x6a, 0x61}, {0x70, 0x6f, 0x6e, 0x65, 0x64, 0x6a, 0x65, 0x6c, 0x6a, 0x61, 0x6b}, {0x75, 0x74, 0x6f, 0x72, 0x61, 0x6b}, {0x73, 0x72, 0x69, 0x6a, 0x65, 0x64, 0x61}, {0xc4, 0x8d, 0x65, 0x74, 0x76, 0x72, 0x74, 0x61, 0x6b}, {0x70, 0x65, 0x74, 0x61, 0x6b}, {0x73, 0x75, 0x62, 0x6f, 0x74, 0x61}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsWide:            [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:        [][]uint8{{0x70, 0x72, 0x2e, 0x20, 0x4b, 0x72, 0x2e}, {0x70, 0x2e, 0x20, 0x4b, 0x72, 0x2e}},
		erasNarrow:             [][]uint8{{0x70, 0x72, 0x2e, 0x6e, 0x2e, 0x65, 0x2e}, {0x41, 0x44}},
		erasWide:               [][]uint8{{0x70, 0x72, 0x69, 0x6a, 0x65, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x61}, {0x70, 0x6f, 0x73, 0x6c, 0x69, 0x6a, 0x65, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x61}},
		timezones:              map[string][]uint8{"OESZ": {0x69, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "SRT": {0x73, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "ARST": {0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "AST": {0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "WAST": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x61, 0x66, 0x72, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "PST": {0x70, 0x61, 0x63, 0x69, 0x66, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "TMT": {0x74, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "SGT": {0x73, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "AKST": {0x61, 0x6c, 0x6a, 0x61, 0xc5, 0xa1, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "AWDT": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "HNT": {0x6e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "CST": {0x73, 0x72, 0x65, 0x64, 0x69, 0xc5, 0xa1, 0x6e, 0x6a, 0x65, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "EST": {0x69, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "MST": {0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x6d, 0x61, 0x6b, 0x61, 0x6f, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "JDT": {0x6a, 0x61, 0x70, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "UYT": {0x75, 0x72, 0x75, 0x67, 0x76, 0x61, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "AKDT": {0x61, 0x6c, 0x6a, 0x61, 0xc5, 0xa1, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "CDT": {0x73, 0x72, 0x65, 0x64, 0x69, 0xc5, 0xa1, 0x6e, 0x6a, 0x65, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "WART": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x2d, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "AEDT": {0x69, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "HADT": {0x68, 0x61, 0x76, 0x61, 0x6a, 0x73, 0x6b, 0x6f, 0x2d, 0x61, 0x6c, 0x65, 0x75, 0x74, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "OEZ": {0x69, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "COT": {0x6b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "SAST": {0x6a, 0x75, 0xc5, 0xbe, 0x6e, 0x6f, 0x61, 0x66, 0x72, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "ART": {0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "ACWDT": {0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x72, 0x65, 0x64, 0x69, 0xc5, 0xa1, 0x6e, 0x6a, 0x65, 0x20, 0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "NZST": {0x6e, 0x6f, 0x76, 0x6f, 0x7a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "VET": {0x76, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "WEZ": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "ADT": {0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "PDT": {0x70, 0x61, 0x63, 0x69, 0x66, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "UYST": {0x75, 0x72, 0x75, 0x67, 0x76, 0x61, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "CLST": {0xc4, 0x8d, 0x69, 0x6c, 0x65, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "EDT": {0x69, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "COST": {0x6b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "EAT": {0x69, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x61, 0x66, 0x72, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "AEST": {0x69, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "CHAST": {0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x61}, "WESZ": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "NZDT": {0x6e, 0x6f, 0x76, 0x6f, 0x7a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "WIB": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "WIT": {0x69, 0x73, 0x74, 0x6f, 0xc4, 0x8d, 0x6e, 0x6f, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "TMST": {0x74, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "∅∅∅": {0x62, 0x72, 0x61, 0x7a, 0x69, 0x6c, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "HKT": {0x68, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0xc5, 0xa1, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "HKST": {0x68, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0xc5, 0xa1, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "BOT": {0x62, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "IST": {0x69, 0x6e, 0x64, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "HAST": {0x68, 0x61, 0x76, 0x61, 0x6a, 0x73, 0x6b, 0x6f, 0x2d, 0x61, 0x6c, 0x65, 0x75, 0x74, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "WAT": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x61, 0x66, 0x72, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "MEZ": {0x73, 0x72, 0x65, 0x64, 0x6e, 0x6a, 0x6f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "JST": {0x6a, 0x61, 0x70, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "BT": {0x62, 0x75, 0x74, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "MYT": {0x6d, 0x61, 0x6c, 0x65, 0x7a, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "ACST": {0x73, 0x72, 0x65, 0x64, 0x6e, 0x6a, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "WITA": {0x73, 0x72, 0x65, 0x64, 0x6e, 0x6a, 0x6f, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x6a, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "CLT": {0xc4, 0x8d, 0x69, 0x6c, 0x65, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "LHDT": {0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65, 0x20, 0x6f, 0x74, 0x6f, 0x6b, 0x61, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "GFT": {0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65, 0x20, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x75, 0x73, 0x6b, 0x65, 0x20, 0x47, 0x76, 0x61, 0x6a, 0x61, 0x6e, 0x65}, "GMT": {0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x7a, 0x61, 0x6c, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "WARST": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x2d, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "ACDT": {0x73, 0x72, 0x65, 0x64, 0x6e, 0x6a, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "CAT": {0x73, 0x72, 0x65, 0x64, 0x6e, 0x6a, 0x6f, 0x61, 0x66, 0x72, 0x69, 0xc4, 0x8d, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "GYT": {0x67, 0x76, 0x61, 0x6a, 0x61, 0x6e, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "MESZ": {0x73, 0x72, 0x65, 0x64, 0x6e, 0x6a, 0x6f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "AWST": {0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "CHADT": {0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x61}, "ChST": {0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x61}, "ECT": {0x65, 0x6b, 0x76, 0x61, 0x64, 0x6f, 0x72, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "HAT": {0x6e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x6f, 0x20, 0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "MDT": {0x6c, 0x6a, 0x65, 0x74, 0x6e, 0x6f, 0x20, 0x6d, 0x61, 0x6b, 0x61, 0x6f, 0x73, 0x6b, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "ACWST": {0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x6f, 0x20, 0x73, 0x72, 0x65, 0x64, 0x69, 0xc5, 0xa1, 0x6e, 0x6a, 0x65, 0x20, 0x7a, 0x61, 0x70, 0x61, 0x64, 0x6e, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65}, "LHST": {0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x20, 0x76, 0x72, 0x69, 0x6a, 0x65, 0x6d, 0x65, 0x20, 0x6f, 0x74, 0x6f, 0x6b, 0x61, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}},
	}
}

// Locale returns the current translators string locale
func (hr *hr_HR) Locale() string {
	return hr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'hr_HR'
func (hr *hr_HR) PluralsCardinal() []locales.PluralRule {
	return hr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'hr_HR'
func (hr *hr_HR) PluralsOrdinal() []locales.PluralRule {
	return hr.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'hr_HR'
func (hr *hr_HR) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod100 := i % 100
	iMod10 := i % 10
	fMod10 := f % 10
	fMod100 := f % 100

	if (v == 0 && iMod10 == 1 && iMod100 != 11) || (fMod10 == 1 && fMod100 != 11) {
		return locales.PluralRuleOne
	} else if (v == 0 && iMod10 >= 2 && iMod10 <= 4 && iMod100 < 12 && iMod100 > 14) || (fMod10 >= 2 && fMod10 <= 4 && fMod100 < 12 && fMod100 > 14) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'hr_HR'
func (hr *hr_HR) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'hr_HR'
func (hr *hr_HR) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := hr.CardinalPluralRule(num1, v1)
	end := hr.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'hr_HR' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hr *hr_HR) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(hr.decimal) + len(hr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, hr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'hr_HR' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (hr *hr_HR) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(hr.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hr.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, hr.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'hr_HR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hr *hr_HR) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hr.currencies[currency]
	l := len(s) + len(hr.decimal) + len(hr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, hr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, hr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, hr.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'hr_HR'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hr *hr_HR) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hr.currencies[currency]
	l := len(s) + len(hr.decimal) + len(hr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, hr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, hr.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, hr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, hr.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, hr.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'hr_HR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hr *hr_HR) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'hr_HR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hr *hr_HR) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, hr.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'hr_HR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hr *hr_HR) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, hr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'hr_HR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hr *hr_HR) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, hr.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, hr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'hr_HR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hr *hr_HR) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'hr_HR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hr *hr_HR) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'hr_HR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hr *hr_HR) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'hr_HR'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hr *hr_HR) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := hr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
