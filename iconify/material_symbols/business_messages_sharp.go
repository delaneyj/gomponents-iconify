package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessMessagesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20V8.75L1 4h21v16H5Zm8.5-2.5l1.4-1.4l-1.1-1.1H18v-2H9l4.5 4.5ZM9 11h9l-4.5-4.5l-1.4 1.4L13.2 9H9v2Z"/>`),
		g.Group(children),
	)
}