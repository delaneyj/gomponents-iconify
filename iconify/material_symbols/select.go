package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Select(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19h2v2Zm-2-4v-2h2v2H3Zm0-4v-2h2v2H3Zm0-4V7h2v2H3Zm0-4q0-.825.588-1.413T5 3v2H3Zm4 16v-2h2v2H7ZM7 5V3h2v2H7Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2q0 .825-.588 1.413T19 21Zm0-4v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V7h2v2h-2Zm0-4V3q.825 0 1.413.588T21 5h-2Z"/>`),
		g.Group(children),
	)
}