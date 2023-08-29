package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Azm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 17l4-4V4h-9L7 8h9v9Zm-6 6l4-4v-9H5l-4 4h9v9Z"/>`),
		g.Group(children),
	)
}