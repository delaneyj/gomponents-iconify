package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatAddOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 9h8V7H7v2Zm0 4h5v-2H7v2Zm10 7v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM3 20V5q0-.825.588-1.413T5 3h12q.825 0 1.413.588T19 5v5.075q-.25-.05-.5-.063T18 10q-2.525 0-4.263 1.75T12 16q0 .25.013.5t.062.5H6l-3 3Z"/>`),
		g.Group(children),
	)
}