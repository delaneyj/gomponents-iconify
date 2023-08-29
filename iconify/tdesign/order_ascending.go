package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrderAscending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2.586L23.414 7L22 8.414l-2-2V20h-2V6.414l-2 2L14.586 7L19 2.586ZM2 4h11v2H2V4Zm0 7h13v2H2v-2Zm0 7h13v2H2v-2Z"/>`),
		g.Group(children),
	)
}