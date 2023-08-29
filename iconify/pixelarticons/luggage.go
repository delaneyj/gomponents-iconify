package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Luggage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 2h6v4h4v14h-2v2h-2v-2H9v2H7v-2H5V6h4V2zm2 4h2V4h-2v2zM7 18h10V8H7v10zm4-8v6H9v-6h2zm4 0v6h-2v-6h2z"/>`),
		g.Group(children),
	)
}