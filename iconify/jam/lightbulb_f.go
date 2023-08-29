package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 14.565H8v-8a1 1 0 1 0-2 0v8H3V13.31a7 7 0 1 1 8 0v1.255zm0 2v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-2h8z"/>`),
		g.Group(children),
	)
}