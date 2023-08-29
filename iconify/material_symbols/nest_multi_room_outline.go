package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestMultiRoomOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21V9l8-6l8 6v12H4Zm2-2h7v-3H6v3Zm9 0h3v-3h-3v3Zm-9-5h3v-2.975H6V14Zm5 0h7v-2.975h-7V14ZM7.3 9.025h9.4L12 5.5L7.3 9.025Z"/>`),
		g.Group(children),
	)
}