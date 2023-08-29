package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoNotTouchOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 18.15l-2-2V4h2v14.15Zm-12-12l-2-2V3h2v3.15Zm4 4l-2-2V1h2v9.15Zm4 2.025h-2V2h2v10.175ZM9.525 21h8.65L9 11.825v5.6l-3.7-2.6L9.525 21Zm-1.05 2L1.2 12.375l1.725-1.65L7 13.575v-3.75L.675 3.5L2.1 2.075l20.3 20.3l-1.425 1.425l-.8-.8h-11.7ZM15 12.175ZM13.6 16.4Z"/>`),
		g.Group(children),
	)
}