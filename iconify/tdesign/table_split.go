package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v2h7V4H4Zm9 0v2h7V4h-7Zm7 4H4v12h16V8Zm-9.002 1.998h2.004v2.004h-2.004V9.998Zm-6 3h2.004v2.004H4.998v-2.004Zm3 0h2.004v2.004H7.998v-2.004Zm3 0h2.004v2.004h-2.004v-2.004Zm3 0h2.004v2.004h-2.004v-2.004Zm3 0h2.004v2.004h-2.004v-2.004Zm-6 3h2.004v2.004h-2.004v-2.004Z"/>`),
		g.Group(children),
	)
}