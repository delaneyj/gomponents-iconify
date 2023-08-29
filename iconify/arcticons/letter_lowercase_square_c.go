package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterLowercaseSquareC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.195 28.567a5.559 5.559 0 0 1-4.829 2.8h0a5.56 5.56 0 0 1-5.56-5.56v-3.614a5.56 5.56 0 0 1 5.56-5.56h0a5.558 5.558 0 0 1 4.823 2.79"/>`),
		g.Group(children),
	)
}