package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 14v4m0 4.01l.01-.011"/><rect width="6" height="12" x="7" y="2" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M3 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H7m3 0h3"/></g>`),
		g.Group(children),
	)
}