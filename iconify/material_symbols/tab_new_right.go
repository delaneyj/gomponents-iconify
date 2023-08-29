package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabNewRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 23l-1.4-1.4l1.575-1.6H15v-2h4.175L17.6 16.4L19 15l4 4l-4 4ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v8.35q-.475-.175-.988-.263T18.976 13q-2.5 0-4.237 1.75T13 19q0 .525.088 1.025t.262.975H5Zm6-4h2v-4h4v-2h-4V7h-2v4H7v2h4v4Z"/>`),
		g.Group(children),
	)
}