package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleUpCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilTriangleUpCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M13 7a.5.5 0 0 1 .429.243l6 10A.5.5 0 0 1 19 18H7a.5.5 0 0 1-.429-.757l6-10A.5.5 0 0 1 13 7ZM7.883 17h10.234L13 8.472L7.883 17Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilTriangleUpCircleFilled0)"/></g>`),
		g.Group(children),
	)
}