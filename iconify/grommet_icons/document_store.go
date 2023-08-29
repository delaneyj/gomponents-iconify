package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 7V1H19.5L23 4.5V23h-6m1-22v5h5M3 12s1-2 6-2s6 2 6 2v9s-1 2-6 2s-6-2-6-2v-9Zm0 5s2 2 6 2s6-2 6-2M3 13s2 2 6 2s6-2 6-2"/>`),
		g.Group(children),
	)
}