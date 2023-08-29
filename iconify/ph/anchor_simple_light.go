package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 114h-24a6 6 0 0 0 0 12h17.8a90.13 90.13 0 0 1-83.8 83.78V93.4a30 30 0 1 0-12 0v116.38A90.13 90.13 0 0 1 38.2 126H56a6 6 0 0 0 0-12H32a6 6 0 0 0-6 6a102 102 0 0 0 204 0a6 6 0 0 0-6-6ZM110 64a18 18 0 1 1 18 18a18 18 0 0 1-18-18Z"/>`),
		g.Group(children),
	)
}