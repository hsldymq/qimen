package component

import (
	"testing"

	"github.com/hsldymq/go-chinese-calendar/sexagenary"
)

func TestNewLeadingHour(t *testing.T) {
	cases := [6][10]string{
		{"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉"},
		{"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未"},
		{"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳"},
		{"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯"},
		{"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑"},
		{"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"},
	}
	expect := [6]HourLeader{
		HourLeaderEnum.JiaZi,
		HourLeaderEnum.JiaXu,
		HourLeaderEnum.JiaShen,
		HourLeaderEnum.JiaWu,
		HourLeaderEnum.JiaChen,
		HourLeaderEnum.JiaYin,
	}
	for idx, hours := range cases {
		for _, each := range hours {
			s, ok := sexagenary.NewSexagenaryTermFromText(each)
			if !ok {
				t.Fatalf("failed to new sexagenary from %s", each)
			}

			lh := NewHourLeader(s)
			if lh != expect[idx] {
				t.Fatalf("leading hour of %s should be %d, got %d", each, expect[idx], lh)
			}
		}
	}

}
