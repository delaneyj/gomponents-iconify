package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stiletto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M946 404q-50 180-50 300v192h-64V640q0-127-42-175q-23 18-49 59.5t-47.5 88t-49.5 98t-55 92t-64 67t-77 26.5q-102 0-201.5-10t-173-33T0 800q0-26 65-79t142-97t113-48q18 18 30.5 28.5t40.5 23t57 12.5q23 0 87-82.5t129.5-189t116.5-215T832 0q80 0 136 65.5t56 158.5q0 54-21 101t-57 79z"/>`),
		g.Group(children),
	)
}