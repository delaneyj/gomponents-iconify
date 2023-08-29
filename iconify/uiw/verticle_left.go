package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.886 2.203a.715.715 0 0 1 1-.012a.69.69 0 0 1 .012.986L8.272 9.855l6.992 6.876a.69.69 0 0 1 0 .986a.715.715 0 0 1-1 0L6.893 10.47v6.839a.695.695 0 0 1-.696.692a.695.695 0 0 1-.697-.692V2.692c0-.382.312-.692.697-.692c.385 0 .697.31.697.692L6.893 9.25Z"/>`),
		g.Group(children),
	)
}