package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.031.031v3.938h3.871L10.031.031zm-3.539 8.99a2.478 2.478 0 0 0-2.049 3.867l-2.385 2.385l.707.707l2.398-2.398c.385.246.838.394 1.328.394a2.479 2.479 0 0 0 2.477-2.477a2.479 2.479 0 0 0-2.476-2.478zm.024 4.078a1.58 1.58 0 0 1-1.578-1.578c0-.87.708-1.578 1.578-1.578a1.58 1.58 0 0 1 1.578 1.578a1.58 1.58 0 0 1-1.578 1.578z"/><path d="M9 5V0H2.042v13.876l1.059-1.059a3.602 3.602 0 0 1-.256-1.316a3.625 3.625 0 1 1 3.625 3.625c-.428 0-.832-.088-1.213-.223L4.222 16H14V5H9z"/></g>`),
		g.Group(children),
	)
}