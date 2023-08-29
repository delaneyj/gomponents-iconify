package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.308 0H1.954A2.121 2.121 0 0 0 0 2.257v-.006V21.75A2.12 2.12 0 0 0 1.948 24h14.356a2.12 2.12 0 0 0 1.954-2.256v.006V2.251a2.114 2.114 0 0 0-1.942-2.25h-.007zM9.131 22.5a1.498 1.498 0 1 1 .002-2.996a1.498 1.498 0 0 1-.002 2.996h-.002z"/>`),
		g.Group(children),
	)
}