package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackHoleTwoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="2"/><path stroke-dasharray="2 2" stroke-linecap="round" d="M10.142 10.363C13.688 6.817 21.914 15.61 16.524 21" opacity=".5"/><path stroke-dasharray="2 2" stroke-linecap="round" d="M13.858 13.637C10.312 17.183 2.086 8.39 7.476 3" opacity=".5"/><path stroke-dasharray="2 2" stroke-linecap="round" d="M10.363 13.858C6.817 10.312 15.61 2.086 21 7.476" opacity=".5"/><path stroke-dasharray="2 2" stroke-linecap="round" d="M13.637 10.142C17.183 13.688 8.39 21.914 3 16.524" opacity=".5"/></g>`),
		g.Group(children),
	)
}