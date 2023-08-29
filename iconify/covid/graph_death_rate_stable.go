package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphDeathRateStable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21.787 22.5h-21v-21"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.162 6.375a4.876 4.876 0 1 0-7.875 3.816v.309a1.5 1.5 0 0 0 1.5 1.5h3a1.5 1.5 0 0 0 1.5-1.5v-.309a4.846 4.846 0 0 0 1.875-3.816ZM11.287 10.5V12m9.535 3.234l2.391 2.391l-2.391 2.391"/><path stroke-linecap="round" stroke-linejoin="round" d="M.787 17.625h.5a10.6 10.6 0 0 0 4.737-1.118L8.446 15.3a6.352 6.352 0 0 1 5.683 0l2.422 1.211a10.6 10.6 0 0 0 4.736 1.118h1.926"/><path d="M13.162 7.125a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m-3.752.75a.375.375 0 1 1 0-.75m0 .75a.375.375 0 1 0 0-.75"/></g>`),
		g.Group(children),
	)
}