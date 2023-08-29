package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RequestQuoteOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 18v.5q0 .2.15.35t.35.15h1q.2 0 .35-.15t.15-.35V18h1q.425 0 .713-.288T15 17v-3q0-.425-.288-.713T14 13h-3v-1h3q.425 0 .713-.288T15 11q0-.425-.288-.713T14 10h-1v-.5q0-.2-.15-.35T12.5 9h-1q-.2 0-.35.15T11 9.5v.5h-1q-.425 0-.713.288T9 11v3q0 .425.288.713T10 15h3v1h-3q-.425 0-.713.288T9 17q0 .425.288.713T10 18h1Zm-5 4q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h9l5 5v13q0 .825-.588 1.413T18 22H6Zm0-2h12V8h-3q-.425 0-.713-.288T14 7V4H6v16ZM6 4v4v-4v16V4Z"/>`),
		g.Group(children),
	)
}