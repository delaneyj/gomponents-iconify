package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tasks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 3h22v4H1V3Zm0 7h22v4H1v-4Zm0 7h22v4H1v-4ZM1 4h15v2H1V4Zm0 7h5v2H1v-2Zm0 7h10v2H1v-2Z"/>`),
		g.Group(children),
	)
}