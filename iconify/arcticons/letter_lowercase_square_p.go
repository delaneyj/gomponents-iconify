package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterLowercaseSquareP(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.44 22.054a5.56 5.56 0 0 0 5.56 5.56h0a5.56 5.56 0 0 0 5.56-5.56V18.44A5.56 5.56 0 0 0 24 12.879h0a5.56 5.56 0 0 0-5.56 5.561m-.001-5.561v22.242"/>`),
		g.Group(children),
	)
}