package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyBitcoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21v-2H6v-2h2V7H6V5h3V3h2v2h2V3h2v2.125q1.3.35 2.15 1.413T18 9q0 .725-.25 1.388t-.7 1.187q.875.525 1.413 1.425T19 15q0 1.65-1.175 2.825T15 19v2h-2v-2h-2v2H9Zm1-10h4q.825 0 1.413-.588T16 9q0-.825-.588-1.413T14 7h-4v4Zm0 6h5q.825 0 1.413-.588T17 15q0-.825-.588-1.413T15 13h-5v4Z"/>`),
		g.Group(children),
	)
}