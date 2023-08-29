package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickThinUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.666 8.333c-.201 0-.404-.045-.596-.141L9.002 5.158L2.931 8.192a1.332 1.332 0 1 1-1.192-2.385l6.666-3.333a1.332 1.332 0 0 1 1.193 0l6.666 3.333a1.332 1.332 0 0 1-.598 2.526z"/><path d="M15.666 12.666a.66.66 0 0 1-.297-.07L9.002 9.412l-6.369 3.184a.667.667 0 1 1-.596-1.193L8.703 8.07a.665.665 0 0 1 .596 0l6.666 3.333a.668.668 0 0 1-.299 1.263z"/></g>`),
		g.Group(children),
	)
}