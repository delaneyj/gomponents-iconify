package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M40 44H8v-4l6-3h20l6 3v4ZM14 19h20"/><path d="M27.74 19L33 37H15l5.26-18"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M27.74 19L33 37H15l5.26-18"/><path stroke="currentColor" stroke-width="4" d="M24 4a8 8 0 0 0-3.876 15h7.752A8 8 0 0 0 24 4Z"/></g>`),
		g.Group(children),
	)
}