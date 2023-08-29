package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoNotTouchSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 18.15l-5.975-5.975H15V2h2v10.175h2V4h2v14.15Zm-8-8l-2-2V1h2v9.15Zm-4-4l-2-2V3h2v3.15ZM8.475 23L1.2 12.375l1.725-1.65L7 13.575v-3.75L.675 3.5L2.1 2.075l20.3 20.3l-1.425 1.425l-.8-.8h-11.7Z"/>`),
		g.Group(children),
	)
}