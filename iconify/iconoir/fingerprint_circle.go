package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FingerprintCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 16v-4.639c0-.51.1-.999.285-1.453M17 16v-3.185m-7.778-5.08A5.506 5.506 0 0 1 12 7c2.28 0 4.203 1.33 4.805 3.15M10 17v-2.177M14 17v-5.147C14 10.83 13.105 10 12 10s-2 .83-2 1.853v.794"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`),
		g.Group(children),
	)
}