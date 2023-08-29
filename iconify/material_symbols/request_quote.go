package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RequestQuote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h8l6 6v12q0 .825-.588 1.413T18 22H6Zm5-3h2v-1h1q.425 0 .713-.288T15 17v-3q0-.425-.288-.713T14 13h-3v-1h4v-2h-2V9h-2v1h-1q-.425 0-.713.288T9 11v3q0 .425.288.713T10 15h3v1H9v2h2v1Zm2.175-11h4l-4-4v4Z"/>`),
		g.Group(children),
	)
}