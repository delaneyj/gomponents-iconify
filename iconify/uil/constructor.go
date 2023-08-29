package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Constructor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 9.67V9.5a7.95 7.95 0 0 0-5.59-7.62h-.06a8.32 8.32 0 0 0-2.59-.36A8.21 8.21 0 0 0 4 9.67a3 3 0 0 0 0 5.66a8 8 0 0 0 8 7.17h.23a8.13 8.13 0 0 0 7.68-7.16A3 3 0 0 0 20 9.67ZM12.18 20.5a6 6 0 0 1-6.09-5h11.77a6.09 6.09 0 0 1-5.68 5Zm6.82-7H5a1 1 0 0 1 0-2h2a1 1 0 0 0 0-2H6a6.4 6.4 0 0 1 3-5.15V7.5a1 1 0 0 0 2 0V3.59a7.34 7.34 0 0 1 .82-.09H12a6.64 6.64 0 0 1 1 .09V7.5a1 1 0 0 0 2 0V4.32a6.65 6.65 0 0 1 1.18.87A6 6 0 0 1 18 9.5h-1a1 1 0 0 0 0 2h2a1 1 0 0 1 0 2Z"/>`),
		g.Group(children),
	)
}