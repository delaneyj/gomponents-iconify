package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellBolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.5 17H4a4 4 0 0 0 2-3v-3a7 7 0 0 1 4-6a2 2 0 1 1 4 0a7 7 0 0 1 4 6v1m-9 5v1a3 3 0 0 0 4.368 2.67M19 16l-2 3h4l-2 3"/>`),
		g.Group(children),
	)
}