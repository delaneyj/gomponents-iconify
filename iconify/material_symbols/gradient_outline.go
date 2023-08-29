package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GradientOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 13v-2h2v2h-2Zm-2 2v-2h2v2H9Zm4 0v-2h2v2h-2Zm2-2v-2h2v2h-2Zm-8 0v-2h2v2H7Zm-2 8q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm2-2h2v-2H7v2Zm4 0h2v-2h-2v2Zm8 0v-2v2ZM5 17h2v-2h2v2h2v-2h2v2h2v-2h2v2h2v-2h-2v-2h2V5H5v8h2v2H5v2Zm0 2V5v14Zm14-6v2v-2Zm-4 4v2h2v-2h-2Z"/>`),
		g.Group(children),
	)
}