package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloatCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 4h6v8H8V4h2zm4 6V6h-4v4h4zM2 6h4v2H2V6zm20 0h-4v2h4V6zm0 4h-4v2h4v-2zM6 10H2v2h4v-2zm-4 4h20v2H2v-2zm20 4H2v2h20v-2z"/>`),
		g.Group(children),
	)
}