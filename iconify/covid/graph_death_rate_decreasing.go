package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphDeathRateDecreasing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21.777 6.375a4.875 4.875 0 1 0-7.877 3.816v.309a1.5 1.5 0 0 0 1.5 1.5h3a1.5 1.5 0 0 0 1.5-1.5v-.309a4.847 4.847 0 0 0 1.877-3.816ZM16.902 10.5V12m4.875 10.5h-21v-21"/><path stroke-linecap="round" stroke-linejoin="round" d="M.777 9.5h1.4a10.826 10.826 0 0 1 8.6 4.25a10.824 10.824 0 0 0 8.6 4.25h3.846"/><path stroke-linecap="round" stroke-linejoin="round" d="M20.833 15.609L23.223 18l-2.39 2.391"/><path d="M18.777 7.112a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m-3.752.75a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		g.Group(children),
	)
}