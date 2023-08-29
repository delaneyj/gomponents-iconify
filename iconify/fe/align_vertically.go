package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVertically(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feAlignVertically0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAlignVertically1" fill="currentColor"><path id="feAlignVertically2" d="M17 13v2a2 2 0 1 1-4 0v-2h-2v4.714C11 18.977 10.105 20 9 20s-2-1.023-2-2.286V13H3a1 1 0 0 1 0-2h4V6.286C7 5.023 7.895 4 9 4s2 1.023 2 2.286V11h2V9a2 2 0 1 1 4 0v2h4a1 1 0 0 1 0 2h-4Z"/></g></g>`),
		g.Group(children),
	)
}