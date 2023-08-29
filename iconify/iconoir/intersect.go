package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intersect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 13.5v3M13.5 21h3m0-12H9.6a.6.6 0 0 0-.6.6v6.9m1.5 4.5h-.9a.6.6 0 0 1-.6-.6v-.9m12 0v.9a.6.6 0 0 1-.6.6h-.9m0-12h.9a.6.6 0 0 1 .6.6v.9m-18 0v-3M7.5 3h3"/><path d="M7.5 15h6.9a.6.6 0 0 0 .6-.6V7.5M4.5 15h-.9a.6.6 0 0 1-.6-.6v-.9m0-9v-.9a.6.6 0 0 1 .6-.6h.9m9 0h.9a.6.6 0 0 1 .6.6v.9"/></g>`),
		g.Group(children),
	)
}