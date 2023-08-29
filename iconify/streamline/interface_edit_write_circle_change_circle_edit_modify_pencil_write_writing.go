package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditWriteCircleChangeCircleEditModifyPencilWriteWriting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 7A6.5 6.5 0 1 1 7 .5"/><path d="m10.5.5l-5 5l-1 4l4-1l5-5"/></g>`),
		g.Group(children),
	)
}