package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M5 19h11a7 7 0 0 0 7-7V9M8 15l-4 4l4 4M19 5H8a7 7 0 0 0-7 7v3M16 1l4 4l-4 4"/>`),
		g.Group(children),
	)
}