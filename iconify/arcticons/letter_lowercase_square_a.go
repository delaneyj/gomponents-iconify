package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterLowercaseSquareA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.56 25.807a5.56 5.56 0 0 1-5.56 5.56h0a5.56 5.56 0 0 1-5.56-5.56v-3.614a5.56 5.56 0 0 1 5.56-5.56h0a5.56 5.56 0 0 1 5.56 5.56m.001 9.175V16.632"/>`),
		g.Group(children),
	)
}