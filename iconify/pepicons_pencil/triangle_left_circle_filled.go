package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeftCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilTriangleLeftCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M7 13a.5.5 0 0 1 .243-.429l10-6A.5.5 0 0 1 18 7v12a.5.5 0 0 1-.757.429l-10-6A.5.5 0 0 1 7 13Zm10 5.117V7.883L8.472 13L17 18.117Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilTriangleLeftCircleFilled0)"/></g>`),
		g.Group(children),
	)
}