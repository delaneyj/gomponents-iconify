package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRightCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilTriangleRightCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M19 13a.5.5 0 0 1-.243.429l-10 6A.5.5 0 0 1 8 19V7a.5.5 0 0 1 .757-.429l10 6A.5.5 0 0 1 19 13ZM9 7.883v10.234L17.528 13L9 7.883Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilTriangleRightCircleFilled0)"/></g>`),
		g.Group(children),
	)
}