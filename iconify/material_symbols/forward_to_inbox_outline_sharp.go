package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardToInboxOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13L4 8v10h9v2H2V4h20v9h-2V8l-8 5Zm0-2l8-5H4l8 5Zm7 12l-1.4-1.4l1.575-1.6H15v-2h4.175l-1.6-1.6L19 15l4 4l-4 4ZM4 8v11v-6v.075V6v2Z"/>`),
		g.Group(children),
	)
}