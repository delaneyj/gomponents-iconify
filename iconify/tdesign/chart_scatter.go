package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartScatter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2h2Zm13 4a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm-6.002-1h2.004v2.004h-2.004V5Zm-5 3h2.004v2.004H5.998V8Zm9 2h2.004v2.004h-2.004V10ZM9 13a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm7.998 2h2.004v2.004h-2.004V15Zm-11 1h2.004v2.004H5.998V16Z"/>`),
		g.Group(children),
	)
}