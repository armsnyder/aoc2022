package main

import (
	"reflect"
	"testing"
)

func TestDay13ParsePacket(t *testing.T) {
	tests := []struct {
		input string
		want  day13Packet
	}{
		{
			input: "[]",
			want:  day13Packet{integer: -1},
		},
		{
			input: "[[1],[2]]",
			want:  day13Packet{integer: -1, list: []day13Packet{{integer: -1, list: []day13Packet{{integer: 1}}}, {integer: -1, list: []day13Packet{{integer: 2}}}}},
		},
		{
			input: "[1,3]",
			want:  day13Packet{integer: -1, list: []day13Packet{{integer: 1}, {integer: 3}}},
		},
		{
			input: "[[],1]",
			want:  day13Packet{integer: -1, list: []day13Packet{{integer: -1}, {integer: 1}}},
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			gotPacket, gotBytesRead := day13ParsePacket([]byte(test.input))
			if !reflect.DeepEqual(test.want, gotPacket) {
				t.Errorf("got %v; want %v", gotPacket, test.want)
			}
			if gotBytesRead != len(test.input) {
				t.Errorf("got %d; want %d", gotBytesRead, len(test.input))
			}
		})
	}
}

func TestDay13Part1(t *testing.T) {
	runDayTests(t, 13, []dayTest{
		// {
		// 	input: "[[],1]\n[[],0]",
		// 	want:  0,
		// },
		{
			input: "[[1],[2,3,4]]\n[[1],4]",
			want:  1,
		},
		{
			input: "[[4,4],4,4]\n[[4,4],4,4,4]",
			want:  1,
		},
		{
			input: "[[[]]]\n[1]",
			want:  1,
		},
		{
			input: "[[],[],[6,0,[3]],[[[],[5,8,6],[8,1]],[]]]\n[[],[2]]",
			want:  1,
		},
		{
			input: `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`,
			want: 13,
		},
	})
}

func TestDay13Part2(t *testing.T) {
	runDayTests(t, 13, []dayTest{
		{
			part2: true,
			input: `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`,
			want: 140,
		},
	})
}
