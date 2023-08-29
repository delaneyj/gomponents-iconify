package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedInboxOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17q-.825 0-1.413-.588T5 15V5q0-.825.588-1.413T7 3h14q.825 0 1.413.588T23 5v10q0 .825-.588 1.413T21 17H7Zm0-5v3h14v-3h-3.55q-.525.9-1.425 1.45T14 14q-1.125 0-2.025-.55T10.55 12H7Zm7 0q.825 0 1.413-.588T16 10h5V5H7v5h5q0 .825.588 1.413T14 12Zm5 9H3q-.825 0-1.413-.588T1 19V7h2v12h16v2ZM7 15h14H7Z"/>`),
		g.Group(children),
	)
}