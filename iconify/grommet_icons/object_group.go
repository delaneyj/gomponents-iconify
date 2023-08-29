package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 1h3v3H1V1Zm19 0h3v3h-3V1ZM4 2h16M4 22h16M1 20h3v3H1v-3Zm19 0h3v3h-3v-3ZM2 4v16M22 4v16M7 7h7v6H7V7Zm10 3v7h-7v-2"/>`),
		g.Group(children),
	)
}