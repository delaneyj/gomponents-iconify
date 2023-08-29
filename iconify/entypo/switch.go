package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Switch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 3H7a7 7 0 1 0 0 14h6a7 7 0 1 0 0-14zm0 12a5 5 0 1 1 .001-10.001A5 5 0 0 1 13 15z"/>`),
		g.Group(children),
	)
}