package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightBranch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M22 8.012c-1.5 0-5.929-.074-7 4.989c-1.083 5.117-5.143 9.847-7 11"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M22 40c-1.5 0-5.929.063-7-5c-1.083-5.116-5.143-9.848-7-11"/><circle cx="8" cy="24" r="4" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 24h14m8 0h12M30 8.001h12m-12 32h12"/></g>`),
		g.Group(children),
	)
}