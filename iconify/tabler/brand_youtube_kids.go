package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandYoutubeKids(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m18.608 17.75l-3.9.268h-.027a13.83 13.83 0 0 0-3.722.828l-2.511.908a4.111 4.111 0 0 1-3.287-.216a3.82 3.82 0 0 1-1.98-2.527l-1.376-6.05a3.669 3.669 0 0 1 .536-2.86A3.964 3.964 0 0 1 4.83 6.44l11.25-2.354c2.137-.448 4.247.85 4.713 2.9l1.403 6.162a3.677 3.677 0 0 1-.697 3.086a4.007 4.007 0 0 1-2.89 1.512v.002z"/><path d="m9 10l1.208 5l4.292-4z"/></g>`),
		g.Group(children),
	)
}