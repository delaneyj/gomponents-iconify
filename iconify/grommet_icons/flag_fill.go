package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" stroke="currentColor" stroke-width="2" d="M2 24V2c8-3.524 11 4.644 20 0v12c-8 4.895-13-4.103-20 0"/>`),
		g.Group(children),
	)
}