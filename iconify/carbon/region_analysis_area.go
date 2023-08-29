package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionAnalysisArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<defs/><path d="M30 23v-2h-2v-2h-2v2h-3v-2h-2v2h-2v2h2v3h-2v2h2v2h2v-2h3v2h2v-2h2v-2h-2v-3zm-4 3h-3v-3h3z" fill="currentColor"/><path d="M16 30a14 14 0 1 1 14-14h-2a12 12 0 1 0-12 12z" fill="currentColor"/>`),
		g.Group(children),
	)
}