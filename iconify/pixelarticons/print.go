package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 2h12v6h4v10h-4v4H6v-4H2V8h4V2zm2 6h8V4H8v4zm-2 8v-4h12v4h2v-6H4v6h2zm2-2v6h8v-6H8z"/>`),
		g.Group(children),
	)
}