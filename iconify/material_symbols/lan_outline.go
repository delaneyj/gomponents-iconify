package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22v-7h3v-4h5V9H8V2h8v7h-3v2h5v4h3v7h-8v-7h3v-2H8v2h3v7H3Zm7-15h4V4h-4v3ZM5 20h4v-3H5v3Zm10 0h4v-3h-4v3ZM12 7ZM9 17Zm6 0Z"/>`),
		g.Group(children),
	)
}