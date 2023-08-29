package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartMinimum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2h2Zm16 1v1c0 2.048-.252 4.71-1.207 6.899c-.482 1.103-1.163 2.133-2.127 2.89c-.979.77-2.199 1.216-3.667 1.215c-1.468-.002-2.687-.449-3.665-1.217C8.37 13.03 7.689 12 7.207 10.899C6.252 8.712 6 6.053 6 4.006v-1h2v1c0 1.939.246 4.276 1.04 6.092c.393.898.899 1.62 1.53 2.116c.615.484 1.396.789 2.43.79c1.034 0 1.816-.304 2.43-.787c.631-.495 1.137-1.219 1.53-2.118C17.754 8.28 18 5.94 18 3.999V3h2ZM5.998 15.998h2.004v2.004H5.998v-2.004Zm3 0h2.004v2.004H8.998v-2.004Zm3 0h2.004v2.004h-2.004v-2.004Zm3 0h2.004v2.004h-2.004v-2.004Zm3 0h2.004v2.004h-2.004v-2.004Z"/>`),
		g.Group(children),
	)
}