package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h18v7h-4v5H4V4zm14 5h2V6h-2v3zm-2-3h-4v2h2v4H8V8h2V6H6v8h10V6zm3 12v2H3v-2h16z"/>`),
		g.Group(children),
	)
}