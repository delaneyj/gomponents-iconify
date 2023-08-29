package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpWidthWide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 4H2v16h20V4zM4 6h2v12H4V6zm16 12h-2V6h2v12z"/>`),
		g.Group(children),
	)
}