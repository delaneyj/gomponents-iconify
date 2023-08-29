package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReceiptTextSendOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 3.5L18 2l-1.5 1.5L15 2l-1.5 1.5L12 2l-1.5 1.5L9 2L7.5 3.5L6 2L4.5 3.5L3 2v20l1.5-1.5L6 22l1.5-1.5L9 22l1.5-1.5L12 22v-2.91H5V4.91h14v8.35l2 1V2l-1.5 1.5M14 23v-4l4-1l-4-1v-4l10 5l-10 5m-2-12v2H6v-2h6m-6 6v-2h6v2H6M18 7v2H6V7h12Z"/>`),
		g.Group(children),
	)
}