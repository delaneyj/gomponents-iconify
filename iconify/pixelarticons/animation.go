package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Animation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2H2v12h2V4h10V2H4zm2 4h12v2H8v10H6V6zm4 4h12v12H10V10zm10 10v-8h-8v8h8z"/>`),
		g.Group(children),
	)
}