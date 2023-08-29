package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventUpcomingOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 22v-2h4V10H5v4H3V4h3V2h2v2h8V2h2v2h3v18h-6Zm-7 2l-1.4-1.4L9.175 20H1v-2h8.175L6.6 15.4L8 14l5 5l-5 5ZM5 8h14V6H5v2Zm0 0V6v2Z"/>`),
		g.Group(children),
	)
}