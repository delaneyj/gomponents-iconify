package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M34 4H14v10h20V4Zm-7 10l6 23H15l6-23m19 30H8v-4l6-3h20l6 3v4ZM12 23.02h24M20.5 4v4m7-4v4"/>`),
		g.Group(children),
	)
}