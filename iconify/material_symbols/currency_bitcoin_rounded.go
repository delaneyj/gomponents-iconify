package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyBitcoinRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21q-.425 0-.713-.288T9 20v-1H7q-.425 0-.713-.288T6 18q0-.425.288-.713T7 17h1V7H7q-.425 0-.713-.288T6 6q0-.425.288-.713T7 5h2V4q0-.425.288-.713T10 3q.425 0 .713.288T11 4v1h2V4q0-.425.288-.713T14 3q.425 0 .713.288T15 4v1.125q1.3.35 2.15 1.413T18 9q0 .725-.25 1.388t-.7 1.187q.875.525 1.413 1.425T19 15q0 1.65-1.175 2.825T15 19v1q0 .425-.287.713T14 21q-.425 0-.713-.288T13 20v-1h-2v1q0 .425-.288.713T10 21Zm0-10h4q.825 0 1.413-.588T16 9q0-.825-.588-1.413T14 7h-4v4Zm0 6h5q.825 0 1.413-.588T17 15q0-.825-.588-1.413T15 13h-5v4Z"/>`),
		g.Group(children),
	)
}