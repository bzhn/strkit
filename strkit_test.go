package strkit

import "testing"

type TestCtolwe struct {
	text   string
	n      int
	ending string
	want   string
}

type TestCtol struct {
	text string
	n    int
	want string
}

type TestFits struct {
	s    string
	n    int
	want bool
}

func TestCtolb(t *testing.T) {
	var cases = []TestCtol{
		{
			text: "Hello, world!",
			n:    7,
			want: "Hello, ",
		}, {
			text: "Huáyŭ 华语 / 華語 or simply Chinese",
			n:    11,
			want: string([]byte{72, 117, 195, 161, 121, 197, 173, 32, 229, 141, 142}),
		}, {
			text: "Hello𒀤World",
			n:    10,
			want: "Hello𒀤W",
		}, {
			text: `Ipsam aspernatur quo suscipit vel
omnis accusamus eius. Consectetur omnis et excepturi libero totam quas aut autem. 
Voluptatem rerum deleniti reiciendis perspiciatis non placeat facilis neque. Quasi aliquam 
porro consequatur. Tempore laborum occaecati sint consequuntur natus consequatur voluptas.`,
			n: 42,
			want: `Ipsam aspernatur quo suscipit vel
omnis ac`,
		},
	}

	for _, tc := range cases {
		res := Ctolb(tc.text, tc.n)
		if res != tc.want {
			t.Logf("ERROR: Want \"%s\", got \"%s\"", tc.want, res)
			t.Fail()
		}
	}
}

func TestCtolc(t *testing.T) {
	var cases = []TestCtol{
		{
			text: "Hello, world!",
			n:    7,
			want: "Hello, ",
		}, {
			text: "Huáyŭ 华语 / 華語 or simply Chinese",
			n:    10,
			want: "Huáyŭ 华语 /",
		}, {
			text: "Hello𒀤World",
			n:    7,
			want: "Hello𒀤W",
		}, {
			text: `Ipsam aspernatur quo suscipit vel
omnis accusamus eius. Consectetur omnis et excepturi libero totam quas aut autem. 
Voluptatem rerum deleniti reiciendis perspiciatis non placeat facilis neque. Quasi aliquam 
porro consequatur. Tempore laborum occaecati sint consequuntur natus consequatur voluptas.`,
			n: 42,
			want: `Ipsam aspernatur quo suscipit vel
omnis ac`,
		},
	}

	for _, tc := range cases {
		res := Ctolc(tc.text, tc.n)
		if res != tc.want {
			t.Logf("ERROR: Want \"%s\", got \"%s\"", tc.want, res)
			t.Fail()
		}
	}
}

func TestCtolweb(t *testing.T) {
	var cases = []TestCtolwe{
		{
			text:   "Hello, world!",
			n:      7,
			ending: "...",
			want:   "Hell...",
		}, {
			text:   "one of the four official languages of Singapore (where it is called Huáyŭ 华语 / 華語 or simply Chinese).",
			n:      88,
			ending: " -> ",
			want:   "one of the four official languages of Singapore (where it is called Huáyŭ 华语 / -> ",
		}, {
			text:   "𒀤",
			n:      6,
			ending: "...",
			want:   string([]byte{240, 146, 128}) + "...",
		}, {
			text: `Ipsam aspernatur quo suscipit vel
omnis accusamus eius. Consectetur omnis et excepturi libero totam quas aut autem. 
Voluptatem rerum deleniti reiciendis perspiciatis non placeat facilis neque. Quasi aliquam 
porro consequatur. Tempore laborum occaecati sint consequuntur natus consequatur voluptas.`,
			n:      42,
			ending: "...",
			want: `Ipsam aspernatur quo suscipit vel
omnis...`,
		},
	}

	for _, tc := range cases {
		res := Ctolweb(tc.text, tc.n, tc.ending)
		if res != tc.want {
			t.Logf("ERROR: Want \"%s\", got \"%s\"", tc.want, res)
			t.Fail()
		}
	}
}

func TestCtolwec(t *testing.T) {
	var cases = []TestCtolwe{
		{
			text:   "Hello, world!",
			n:      7,
			ending: "...",
			want:   "Hell...",
		}, {
			text:   "one of the four official languages of Singapore (where it is called Huáyŭ 华语 / 華語 or simply Chinese).",
			n:      88,
			ending: " -> ",
			want:   "one of the four official languages of Singapore (where it is called Huáyŭ 华语 / 華語 or -> ",
		}, {
			text:   "𒀤",
			n:      4,
			ending: "...",
			want:   "𒀤...",
		}, {
			text: `Ipsam aspernatur quo suscipit vel
omnis accusamus eius. Consectetur omnis et excepturi libero totam quas aut autem. 
Voluptatem rerum deleniti reiciendis perspiciatis non placeat facilis neque. Quasi aliquam 
porro consequatur. Tempore laborum occaecati sint consequuntur natus consequatur voluptas.`,
			n:      42,
			ending: "...",
			want: `Ipsam aspernatur quo suscipit vel
omnis...`,
		},
	}

	for _, tc := range cases {
		res := Ctolwec(tc.text, tc.n, tc.ending)
		if res != tc.want {
			t.Logf("ERROR: Want \"%s\", got \"%s\"", tc.want, res)
			t.Fail()
		}
	}
}

func TestLenByte(t *testing.T) {
	type test struct {
		text string
		want int
	}

	var cases = []test{{
		text: "hello",
		want: 5,
	}, {
		text: "華語",
		want: 6,
	}, {
		text: "𒀤",
		want: 4,
	}}

	for _, tc := range cases {
		res := LenByte(tc.text)
		if res != tc.want {
			t.Logf("ERROR: Want \"%d\", got \"%d\". String: %s", tc.want, res, tc.text)
			t.Fail()
		}
	}
}

func TestLenChar(t *testing.T) {
	type test struct {
		text string
		want int
	}

	var cases = []test{{
		text: "hello",
		want: 5,
	}, {
		text: "華語",
		want: 2,
	}, {
		text: "𒀤",
		want: 1,
	}}

	for _, tc := range cases {
		res := LenChar(tc.text)
		if res != tc.want {
			t.Logf("ERROR: Want \"%d\", got \"%d\". String: %s", tc.want, res, tc.text)
			t.Fail()
		}
	}
}

func TestFitsC(t *testing.T) {
	var cases = []TestFits{
		{
			s:    "Should fit",
			n:    10,
			want: true,
		}, {
			s:    "華語",
			n:    2,
			want: true,
		}, {
			s:    "Too much charachters",
			n:    10,
			want: false,
		}, {
			s:    "𒀤",
			n:    2,
			want: true,
		}, {
			s: `Multi
line
string`,
			n:    17,
			want: true,
		},
	}

	for _, tc := range cases {
		res := Fitsc(tc.s, tc.n)
		if res != tc.want {
			t.Logf("ERROR: Want \"%t\", got \"%t\". String: %s, len: %d", tc.want, res, tc.s, tc.n)
			t.Fail()
		}
	}
}

func TestFitsB(t *testing.T) {
	var cases = []TestFits{
		{
			s:    "Should fit",
			n:    10,
			want: true,
		}, {
			s:    "華語",
			n:    2,
			want: false,
		}, {
			s:    "Too much charachters",
			n:    10,
			want: false,
		}, {
			s:    "𒀤",
			n:    2,
			want: false,
		}, {
			s: `Multi
line
string`,
			n:    17,
			want: true,
		},
	}

	for _, tc := range cases {
		res := Fitsb(tc.s, tc.n)
		if res != tc.want {
			t.Logf("ERROR: Want \"%t\", got \"%t\". String: %s, len: %d", tc.want, res, tc.s, tc.n)
			t.Fail()
		}
	}
}
