package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThinDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.031 15.97L2.339 9.665c-.459-.459-.296-1.16-.094-1.359c.403-.402.998-.342 1.402.061l4.438 3.896L8.086.948a.94.94 0 0 1 .952-.927a.937.937 0 0 1 .95.927l.013 11.212l4.271-3.854a1.036 1.036 0 0 1 1.461 0c.405.4.304.95-.101 1.352L9.031 15.97z"/>`),
		g.Group(children),
	)
}