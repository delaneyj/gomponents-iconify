package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeftFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.48 11.864a1 1 0 0 1 0-1.728l9.016-5.259A1 1 0 0 1 17 5.741V16.26a1 1 0 0 1-1.504.864l-9.015-5.26Z" opacity=".2"/><path d="M4 10a.5.5 0 0 1 .243-.429l10-6A.5.5 0 0 1 15 4v12a.5.5 0 0 1-.757.429l-10-6A.5.5 0 0 1 4 10Zm10 5.117V4.883L5.472 10L14 15.117Z"/></g>`),
		g.Group(children),
	)
}