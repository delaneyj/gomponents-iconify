package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7v11m0-4h11m4 0h3m0 4v-8a2 2 0 0 0-2-2h-7m-1 3v3m-5-4a1 1 0 1 0 2 0a1 1 0 1 0-2 0M3 3l18 18"/>`),
		g.Group(children),
	)
}