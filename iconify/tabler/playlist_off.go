package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 14a3 3 0 1 0 3 3m0-4V4h4m-8 1H9M5 5H3m0 4h6m0 4H3M3 3l18 18"/>`),
		g.Group(children),
	)
}