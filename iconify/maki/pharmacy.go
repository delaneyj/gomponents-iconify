package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pharmacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m9.5 4l1.07-1.54c.06.005.12.005.18 0A1.25 1.25 0 1 0 9.5 1.25v.1L7 4h2.5zM12 6V5H3v1l1.5 3.5L3 13v1h9v-1l-1-3.5L12 6zm-2 4H8v2H7v-2H5V9h2V7h1v2h2v1z"/>`),
		g.Group(children),
	)
}