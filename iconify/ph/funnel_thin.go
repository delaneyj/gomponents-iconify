package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M227 51.15A11.85 11.85 0 0 0 216 44H40a12 12 0 0 0-8.88 20.07l.05.05l67.73 72.31a4 4 0 0 1 1.08 2.74V216a12 12 0 0 0 18.66 10l32-21.33a12 12 0 0 0 5.35-10v-55.5a4 4 0 0 1 1.08-2.74l67.78-72.36A11.85 11.85 0 0 0 227 51.15Zm-8 7.5L151.24 131a12 12 0 0 0-3.24 8.21v55.49a4 4 0 0 1-1.78 3.33l-32 21.33A4 4 0 0 1 108 216v-76.83a12 12 0 0 0-3.24-8.21L37.05 58.67A4 4 0 0 1 40 52h176a4 4 0 0 1 3 6.65Z"/>`),
		g.Group(children),
	)
}