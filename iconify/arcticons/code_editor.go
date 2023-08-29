package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeEditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.592 37.092L4.5 24l13.092-13.092m2.824 26.202l7.218-26.219m2.774.017L43.5 24L30.408 37.092"/>`),
		g.Group(children),
	)
}