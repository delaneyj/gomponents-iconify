package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationSearching(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 23v-2q-3.125-.35-5.363-2.588T3.05 13.05h-2v-2h2q.35-3.125 2.588-5.363T11 3.1v-2h2v2q3.125.35 5.363 2.588t2.587 5.362h2v2h-2q-.35 3.125-2.587 5.363T13 21v2h-2Zm1-3.95q2.9 0 4.95-2.05T19 12.05q0-2.9-2.05-4.95T12 5.05q-2.9 0-4.95 2.05T5 12.05q0 2.9 2.05 4.95T12 19.05Z"/>`),
		g.Group(children),
	)
}