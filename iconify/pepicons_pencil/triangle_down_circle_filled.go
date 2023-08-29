package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDownCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilTriangleDownCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M13 19a.5.5 0 0 1-.429-.243l-6-10A.5.5 0 0 1 7 8h12a.5.5 0 0 1 .429.757l-6 10A.5.5 0 0 1 13 19Zm5.117-10H7.883L13 17.528L18.117 9Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilTriangleDownCircleFilled0)"/></g>`),
		g.Group(children),
	)
}