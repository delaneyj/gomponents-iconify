package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Search(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m15 15l7 7l-7-7Zm-5.5 2a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Z"/>`),
		g.Group(children),
	)
}