package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chats(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13.584 7a5.001 5.001 0 1 0-8.809 4.675L4 14l2.325-.775c.214.136.44.256.675.359"/><path fill="currentColor" d="M15 20a5 5 0 1 1 4.225-2.325L20 20l-2.325-.775A4.976 4.976 0 0 1 15 20z"/></g>`),
		g.Group(children),
	)
}