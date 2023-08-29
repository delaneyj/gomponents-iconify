package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBoldCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8.75 11.25H12a2.25 2.25 0 0 0 0-4.5H9.522a.772.772 0 0 0-.772.772v3.728Zm0 1.5v4.073c0 .236.19.427.426.427H13a2.25 2.25 0 0 0 0-4.5H8.75Z"/><path fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm5.25-4.478A2.272 2.272 0 0 1 9.522 5.25H12a3.75 3.75 0 0 1 2.665 6.389A3.75 3.75 0 0 1 13 18.75H9.176a1.926 1.926 0 0 1-1.926-1.927V7.522Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}