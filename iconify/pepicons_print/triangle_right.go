package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M14.514 11.25L8 7.472v7.556l6.514-3.778Zm2.495.865a1 1 0 0 0 0-1.73L7.502 4.871A1 1 0 0 0 6 5.736v11.028a1 1 0 0 0 1.502.865l9.507-5.514Z" opacity=".2"/><path d="M16 10a.5.5 0 0 1-.243.429l-10 6A.5.5 0 0 1 5 16V4a.5.5 0 0 1 .757-.429l10 6A.5.5 0 0 1 16 10ZM6 4.883v10.234L14.528 10L6 4.883Z"/></g>`),
		g.Group(children),
	)
}