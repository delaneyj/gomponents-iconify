package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestMultiRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 9H4l8-6l8 6ZM4 21h9v-4H4v4Zm11 0h5v-4h-5v4ZM4 15h5v-4H4v4Zm7 0h9v-4h-9v4Z"/>`),
		g.Group(children),
	)
}