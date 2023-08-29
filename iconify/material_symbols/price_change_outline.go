package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PriceChangeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17h2v-1h1q.425 0 .713-.288T12 15v-3q0-.425-.288-.713T11 11H8v-1h4V8h-2V7H8v1H7q-.425 0-.713.288T6 9v3q0 .425.288.713T7 13h3v1H6v2h2v1Zm8-.75l2-2h-4l2 2ZM14 10h4l-2-2l-2 2ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}