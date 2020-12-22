package link

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		desc string
		name string
		want []*Link
	}{
		{
			desc: "ex1.html",
			name: "ex1.html",
			want: []*Link{
				&Link{HREF: "/other-page", Text: "A link to another page"},
			},
		},
		{
			desc: "ex2.html",
			name: "ex2.html",
			want: []*Link{
				&Link{
					HREF: "https://www.twitter.com/joncalhoun",
					Text: "Check me out on twitter",
				},
				&Link{
					HREF: "https://github.com/gophercises",
					Text: "Gophercises is on Github !",
				},
			},
		},
		{
			desc: "ex3.html",
			name: "ex3.html",
			want: []*Link{
				&Link{
					HREF: "#",
					Text: "Login",
				},
				&Link{
					HREF: "/lost",
					Text: "Lost? Need help?",
				},
				&Link{
					HREF: "https://twitter.com/marcusolsson",
					Text: "@marcusolsson",
				},
			},
		},
		{
			desc: "ex4.html",
			name: "ex4.html",
			want: []*Link{
				&Link{HREF: "/dog-cat", Text: "dog cat"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			f, err := os.Open(tc.name)
			if err != nil {
				t.Fatalf("Parse() unable to read file %s. %v", tc.name, err)
			}

			got, err := Parse(f)
			if err != nil {
				t.Fatalf("Parse() unable to parse contents of %s. %v", tc.name, err)
			}

			if len(got) != len(tc.want) {
				t.Fatalf("Parse()=%v, want=%v", got, tc.want)
			}

			for i := range got {
				if *got[i] != *tc.want[i] {
					t.Errorf("Parse() result %d got:%v, want:%v", i, *got[i], *tc.want[i])
				}
			}
		})
	}
}
