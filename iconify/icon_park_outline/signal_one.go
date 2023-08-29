package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M38.142 38.142A19.937 19.937 0 0 0 44 24a19.937 19.937 0 0 0-5.858-14.142m-28.284 0A19.937 19.937 0 0 0 4 24a19.938 19.938 0 0 0 5.858 14.142M15.1 14.1A13.956 13.956 0 0 0 11 24a13.96 13.96 0 0 0 4.1 9.9m17.8 0A13.956 13.956 0 0 0 37 24a13.96 13.96 0 0 0-4.1-9.9"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 44V34m0-30v10"/><path fill="currentColor" d="M24 27a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`),
		g.Group(children),
	)
}