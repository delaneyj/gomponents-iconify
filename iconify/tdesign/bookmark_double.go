package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23 0l.003 18.419L21 16.415V2.001l-10.999.001v-2L23 .001ZM3 4h16v19.943l-8-5.714l-8 5.714V4Zm2 2v14.057l6-4.286l6 4.286V6H5Z"/>`),
		g.Group(children),
	)
}