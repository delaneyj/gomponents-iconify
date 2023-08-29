package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AntiCorrosion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M24 41.996c13.333.103 20-.989 20-3.275c0-3.428-15.586-20.718-20-20.718S4 35.65 4 38.72c0 2.048 6.667 3.139 20 3.275Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="m21.06 29.661l-3.62 4.34"/><path fill="currentColor" d="M24.5 11a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm9.5 5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-22 5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm5.5-6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/></g>`),
		g.Group(children),
	)
}