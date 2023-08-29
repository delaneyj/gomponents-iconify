package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Helicopter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 8V6H8v2h10v4H4V8H2v8h2v-2h6.22l2.053 8.213A4.992 4.992 0 0 0 17.123 26H26a4.005 4.005 0 0 0 4-4v-2.638a2 2 0 0 0-.464-1.28l-4.468-5.362a1.997 1.997 0 0 0-1.536-.72H20V8Zm-4 16h-8.877a2.995 2.995 0 0 1-2.91-2.272L12.281 14H18v6h10v2a2.002 2.002 0 0 1-2 2Zm-2.468-10l3.333 4H20v-4Z"/>`),
		g.Group(children),
	)
}