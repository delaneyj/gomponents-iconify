package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatQuoteRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.175 17q-.75 0-1.15-.638t-.05-1.312L17 13h-2q-.825 0-1.413-.588T13 11V8q0-.825.588-1.413T15 6h3q.825 0 1.413.588T20 8v4.525q0 .225-.038.463t-.162.437l-1.425 2.825q-.175.35-.5.55t-.7.2Zm-9 0q-.75 0-1.15-.638t-.05-1.312L8 13H6q-.825 0-1.413-.588T4 11V8q0-.825.588-1.413T6 6h3q.825 0 1.413.588T11 8v4.525q0 .225-.038.463t-.162.437L9.375 16.25q-.175.35-.5.55t-.7.2Z"/>`),
		g.Group(children),
	)
}