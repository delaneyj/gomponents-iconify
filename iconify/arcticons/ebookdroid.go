package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ebookdroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".87" d="M36.48 24H6.64M24.9 9.96H4.5m20.4 23.25V10h11.58a7 7 0 1 1 0 14h0a7 7 0 1 1 0 14H4.5"/>`),
		g.Group(children),
	)
}