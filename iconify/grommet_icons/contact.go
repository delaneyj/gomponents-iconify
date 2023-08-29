package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Contact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 2h21v16h-8l-8 4v-4H1V2Zm5 8h1v1H6v-1Zm5 0h1v1h-1v-1Zm5 0h1v1h-1v-1Z"/>`),
		g.Group(children),
	)
}