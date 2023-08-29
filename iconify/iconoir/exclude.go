package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exclude(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.5 15h-.9a.6.6 0 0 0-.6.6v4.8a.6.6 0 0 0 .6.6h10.8a.6.6 0 0 0 .6-.6V9.6a.6.6 0 0 0-.6-.6h-4.8a.6.6 0 0 0-.6.6v.9"/><path d="M13.5 15h.9a.6.6 0 0 0 .6-.6v-.9m-6 0v.9a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-.9"/><path d="M9 10.5v-.9a.6.6 0 0 1 .6-.6h.9"/></g>`),
		g.Group(children),
	)
}