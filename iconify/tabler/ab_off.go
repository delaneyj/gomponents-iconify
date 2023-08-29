package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ABOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 16v-5.5a2.5 2.5 0 0 1 5 0V16m0-4H3m9 0v6m0-12v2m4 0h3a2 2 0 1 1 0 4h-3m3 0a2 2 0 0 1 .83 3.82M16 12V8M3 3l18 18"/>`),
		g.Group(children),
	)
}