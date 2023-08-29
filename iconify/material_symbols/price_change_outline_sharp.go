package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PriceChangeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm2-2h16V6H4v12Zm0 0V6v12Zm4-1h2v-1h1q.425 0 .713-.288T12 15v-3q0-.425-.288-.713T11 11H8v-1h4V8h-2V7H8v1H7q-.425 0-.713.288T6 9v3q0 .425.288.713T7 13h3v1H6v2h2v1Zm8-.75l2-2h-4l2 2ZM14 10h4l-2-2l-2 2Z"/>`),
		g.Group(children),
	)
}