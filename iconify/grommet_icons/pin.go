package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m16 3l-6 6s-4-1-7 2l10 10c3-3 2-7 2-7l6-6l-5-5ZM1 23l7-7m6-15l9 9"/>`),
		g.Group(children),
	)
}