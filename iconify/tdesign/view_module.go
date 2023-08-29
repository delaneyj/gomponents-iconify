package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewModule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1V3Zm2 2v6h4.667V5H3Zm6.667 0v6h4.666V5H9.667Zm6.666 0v6H21V5h-4.667ZM21 13h-4.667v6H21v-6Zm-6.667 6v-6H9.667v6h4.666Zm-6.666 0v-6H3v6h4.667Z"/>`),
		g.Group(children),
	)
}