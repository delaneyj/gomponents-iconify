package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xingtu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M15.556 39L8.177 11.017a.458.458 0 0 1 .314-.561a.386.386 0 0 1 .374.062l33.608 27.67a.472.472 0 0 1 .064.624a.343.343 0 0 1-.311.188h-26.67ZM44.98 9.581a.458.458 0 0 0-.507-.575a.42.42 0 0 0-.18.076l-18.117 12.49l12.995 10.68L44.98 9.58ZM2.188 38.189a.483.483 0 0 0-.124.623a.441.441 0 0 0 .374.188h10.621l-1.812-7.058l-9.06 6.247Z"/>`),
		g.Group(children),
	)
}