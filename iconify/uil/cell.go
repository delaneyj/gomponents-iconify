package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.49 4.73L17 2.17a1 1 0 0 0-1 0l-4 2.28l-4-2.28a1 1 0 0 0-1 0L2.51 4.73A1 1 0 0 0 2 5.6v5.12a1 1 0 0 0 .51.87l4 2.27v4.54a1 1 0 0 0 .51.87l4.5 2.56a1 1 0 0 0 1 0L17 19.27a1 1 0 0 0 .51-.87v-4.54l4-2.27a1 1 0 0 0 .51-.87V5.6a1 1 0 0 0-.53-.87ZM4 10.14v-4l3.5-2l3.5 2v4l-3.5 2Zm11.5 7.68l-3.5 2l-3.5-2v-4l3.5-2l3.5 2Zm4.5-7.68l-3.5 2l-3.5-2v-4l3.5-2l3.5 2Z"/>`),
		g.Group(children),
	)
}