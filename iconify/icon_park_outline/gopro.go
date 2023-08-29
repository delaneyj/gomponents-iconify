package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gopro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 7h13v34H5z"/><rect width="19" height="22" x="24" y="13" stroke="currentColor" stroke-width="4" rx="3"/><circle cx="33.5" cy="24.5" r="3.5" stroke="currentColor" stroke-width="4"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 24h6"/><rect width="5" height="5" x="9" y="15" fill="currentColor" rx="2.5"/></g>`),
		g.Group(children),
	)
}