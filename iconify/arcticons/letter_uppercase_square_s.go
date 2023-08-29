package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterUppercaseSquareS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.71 10.5A6.74 6.74 0 0 0 15 17.25h0A6.74 6.74 0 0 0 21.71 24H24m0 0h2.29A6.74 6.74 0 0 1 33 30.75h0a6.74 6.74 0 0 1-6.73 6.75m6.11-24.72c-1.86-1.56-3.87-2.28-8.38-2.28h-2.29"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.62 35.22c1.86 1.56 3.87 2.28 8.38 2.28h2.29"/>`),
		g.Group(children),
	)
}