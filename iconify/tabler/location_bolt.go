package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationBolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.05 20.1L10 14l-7-3.5a.55.55 0 0 1 0-1L21 3l-3.312 9.173M19 16l-2 3h4l-2 3"/>`),
		g.Group(children),
	)
}