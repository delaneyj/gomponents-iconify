package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditQuillChangeEditFeatherModifyQuillWriteWriting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.51 10.05a10.21 10.21 0 0 0 3 .45A10 10 0 0 0 13.5 1a10.16 10.16 0 0 0-3-.45a10 10 0 0 0-9.99 9.5Z"/><path d="M1 13.5a9.65 9.65 0 0 1-.49-3.45M8.01 5l2.57 2.57"/></g>`),
		g.Group(children),
	)
}