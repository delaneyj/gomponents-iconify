package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Domain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M13 3v4v-4ZM9 3v4v-4ZM5 3v4v-4ZM1 7h22H1Zm0 14h22V3H1v18Z"/>`),
		g.Group(children),
	)
}