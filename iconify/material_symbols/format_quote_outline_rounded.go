package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatQuoteOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 11h3V8h-3v3Zm-9 0h3V8H6v3Zm11.175 6q-.75 0-1.15-.638t-.05-1.312L17 13h-2q-.825 0-1.413-.588T13 11V8q0-.825.588-1.413T15 6h3q.825 0 1.413.588T20 8v4.525q0 .225-.038.463t-.162.437l-1.425 2.825q-.175.35-.5.55t-.7.2Zm-9 0q-.75 0-1.15-.638t-.05-1.312L8 13H6q-.825 0-1.413-.588T4 11V8q0-.825.588-1.413T6 6h3q.825 0 1.413.588T11 8v4.525q0 .225-.038.463t-.162.437L9.375 16.25q-.175.35-.5.55t-.7.2ZM7.5 9.5Zm9 0Z"/>`),
		g.Group(children),
	)
}