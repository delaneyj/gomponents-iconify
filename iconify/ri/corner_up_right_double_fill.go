package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpRightDoubleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19v-9h5.586V4.586L16 11l-6.414 6.414V12H6v7H4Zm9.836-12.95l1.414-1.414L21.614 11l-6.364 6.364l-1.414-1.414l4.95-4.95l-4.95-4.95Z"/>`),
		g.Group(children),
	)
}