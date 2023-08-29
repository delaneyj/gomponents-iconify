package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScooterLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 138a34.32 34.32 0 0 0-4.89.36l-11.27-33.82l-22.15-66.44A6 6 0 0 0 168 34h-32a6 6 0 0 0 0 12h27.68l19.75 59.25L133.07 170H77.94a34 34 0 1 0-1.44 12H136a6 6 0 0 0 4.74-2.32L188 118.93l7.74 23.23A34 34 0 1 0 212 138ZM44 194a22 22 0 1 1 22-22a22 22 0 0 1-22 22Zm168 0a22 22 0 1 1 22-22a22 22 0 0 1-22 22Z"/>`),
		g.Group(children),
	)
}