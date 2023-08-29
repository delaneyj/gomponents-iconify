package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Directions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m17 11l5-5l-5-5m5 5h-4a6 6 0 0 0-6 6v12M7 6l-5 5l5 5m-5-5h4a6 6 0 0 1 6 6v7"/>`),
		g.Group(children),
	)
}