package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 2H6v4H2v2h14v14h2v-4h4v-2h-4V6H8V2zm0 8H6v8h8v-2H8v-6z"/>`),
		g.Group(children),
	)
}