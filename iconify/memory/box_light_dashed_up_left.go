package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLightDashedUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 22v-2h2v2h-2m0-4v-3h2v3h-2m0-5v-3h3v2h-1v1h-2m12-1h-2v-2h2v2m-4 0h-3v-2h3v2Z"/>`),
		g.Group(children),
	)
}