package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForestOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18H1.825q-.6 0-.875-.525T1 16.45L3.85 12q-.6 0-.85-.537t.1-1.038l5.075-7.25q.3-.425.825-.425t.825.425L12 6.3l2.175-3.125q.3-.425.825-.425t.825.425l5.075 7.25q.35.5.1 1.038t-.85.537L23 16.45q.325.5.05 1.025t-.875.525H17v2q0 .825-.588 1.413T15 22q-.825 0-1.413-.588T13 20v-2h-2v2q0 .825-.588 1.413T9 22q-.825 0-1.413-.588T7 20v-2Zm9.725-2h3.625l-3.875-6h1.675L15 5.5l-1.775 2.525l1.675 2.4q.35.5.1 1.038t-.85.537l2.575 4ZM3.65 16h10.7l-3.875-6h1.675L9 5.5L5.85 10h1.675L3.65 16Zm0 0h3.875H5.85h6.3h-1.675h3.875h-10.7Zm13.075 0H14.15H15h-1.775h4.925h-1.675h3.875h-3.625ZM13 18h4h-4Zm3.175 0Z"/>`),
		g.Group(children),
	)
}