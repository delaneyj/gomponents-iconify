package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M13.8 13.8L18 18l-4.2-4.2ZM10.5 15a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Z"/>`),
		g.Group(children),
	)
}