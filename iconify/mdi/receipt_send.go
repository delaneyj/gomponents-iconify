package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReceiptSend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 20.5L3 22V2l1.5 1.5L6 2l1.5 1.5L9 2l1.5 1.5L12 2l1.5 1.5L15 2l1.5 1.5L18 2l1.5 1.5L21 2v12.26l-9-4.5V22l-1.5-1.5L9 22l-1.5-1.5L6 22l-1.5-1.5M14 19l4-1l-4-1v-4l10 5l-10 5v-4Z"/>`),
		g.Group(children),
	)
}