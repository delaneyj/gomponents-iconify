package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimalisticMagniferBugLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="11.5" cy="11.5" r="9.5" opacity=".5"/><path stroke-linecap="round" d="M11.5 15.5a3 3 0 0 1-3-3v-2m3 5a3 3 0 0 0 3-3v-2m-3 5v-5m3 0a3 3 0 1 0-6 0m6 0h-6m6.072 1H16m-9 0h1.5m6 2.5l1 .5m-7-.5l-1 .5m7-5.5l1-.5m-7 .5l-1-.5M20 20l2 2"/></g>`),
		g.Group(children),
	)
}