package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 3h22v18H1V3Zm5 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm17 6l-5-6l-6 7l-3-3l-8 8"/>`),
		g.Group(children),
	)
}