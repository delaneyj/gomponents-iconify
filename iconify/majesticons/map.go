package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M21.472 3.118A1 1 0 0 1 22 4v12a1 1 0 0 1-.445.832L16 20.535V6.131l4.445-2.963a1 1 0 0 1 1.027-.05zM14 6.131l-4-2.666v14.404l4 2.666V6.131zM8 17.87l-4.445 2.963A1 1 0 0 1 2 20V8a1 1 0 0 1 .445-.832L8 3.465v14.404z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}