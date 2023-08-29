package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Materialislands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.45 30.24L24 28.18h0l-13.45 2.06L24 4.5l13.45 25.74z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.55 30.24L24 28.18h0l13.45 2.06L24 43.5L10.55 30.24zM24 4.5v39"/>`),
		g.Group(children),
	)
}