package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowAlignV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 11v2h10v-2H7Zm.757 7l1.415-1.414L11 18.414V14h2v4.414l1.828-1.828L16.243 18L12 22.243L7.757 18Zm8.486-12l-1.414 1.414L13 5.586V10h-2V5.586L9.172 7.414L7.757 6L12 1.757L16.243 6Z"/>`),
		g.Group(children),
	)
}