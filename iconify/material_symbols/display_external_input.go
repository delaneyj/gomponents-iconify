package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisplayExternalInput(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 22l-1.425-1.425l1.6-1.575H14v-2h4.175L16.6 15.4L18 14l4 4l-4 4ZM5 21q-.825 0-1.413-.588T3 19v-4h2v4h4v2H5ZM3 9V5q0-.825.588-1.413T5 3h4v2H5v4H3Zm16 0V5h-4V3h4q.825 0 1.413.588T21 5v4h-2Z"/>`),
		g.Group(children),
	)
}