package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterUppercaseSquareB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-25.52 32v-27"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15 10.5h11.29A6.74 6.74 0 0 1 33 17.25h0A6.74 6.74 0 0 1 26.29 24H15m0 0h11.29A6.74 6.74 0 0 1 33 30.75h0a6.74 6.74 0 0 1-6.73 6.75H15"/>`),
		g.Group(children),
	)
}