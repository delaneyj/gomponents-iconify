package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm7-3.5v1a.5.5 0 0 0 1 0v-1h2v1a2.5 2.5 0 0 1-5 0v-1h2Zm7 0v1a.5.5 0 0 0 1 0v-1h2v1a2.5 2.5 0 0 1-5 0v-1h2ZM11 14h2.004v2.004H11V14Z"/>`),
		g.Group(children),
	)
}