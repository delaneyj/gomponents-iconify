package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditFlipHorizontalOneArrowDesignFlipReflect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 3l4 4l-4 4V3zm13 0l-4 4l4 4V3zM7 .5v1m0 2v1m0 2v1m0 2v1m0 2v1"/>`),
		g.Group(children),
	)
}