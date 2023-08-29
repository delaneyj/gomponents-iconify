package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Desktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 1h22v18H1V1Zm4 22h14H5Zm5-4v4v-4Zm4 0v4v-4Z"/>`),
		g.Group(children),
	)
}