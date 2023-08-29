package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 12c0-3.31-2.69-6-6-6s-6 2.69-6 6s2.69 6 6 6s6-2.69 6-6m-6-4a3.999 3.999 0 1 1 .002 8.002A3.999 3.999 0 0 1 12 8m8-4H4v16h16V4m2-2v20H2V2h20Z"/>`),
		g.Group(children),
	)
}