package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warehouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13.5 5a.5.5 0 0 1-.22-.05L7.5 2L1.72 4.93A.514.514 0 0 1 1.28 4L7.5.92L13.72 4a.512.512 0 0 1-.22 1zM5 10H2v3h3v-3zm4 0H6v3h3v-3zm4 0h-3v3h3v-3zm-2-4H8v3h3V6zM7 6H4v3h3V6z"/>`),
		g.Group(children),
	)
}