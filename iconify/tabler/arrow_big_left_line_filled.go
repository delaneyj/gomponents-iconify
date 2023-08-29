package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBigLeftLineFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M9.586 4L3 10.586a2 2 0 0 0 0 2.828L9.586 20a2 2 0 0 0 2.18.434l.145-.068A2 2 0 0 0 13 18.586V16h5a1 1 0 0 0 1-1V9l-.007-.117A1 1 0 0 0 18 8l-5-.001V5.414A2 2 0 0 0 9.586 4z"/><path fill="currentColor" d="M4.415 12L11 5.414V9l.007.117A1 1 0 0 0 12 10l5-.001v4L12 14a1 1 0 0 0-1 1v3.586L4.415 12zM21 8a1 1 0 0 1 .993.883L22 9v6a1 1 0 0 1-1.993.117L20 15V9a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}