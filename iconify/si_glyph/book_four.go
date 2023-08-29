package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 0v16H3.016v-2H4v-2h-.984V4H4V2h-.984V0H15zM2 2v2h1V2H2zm0 10v2h1v-2H2zm15-1.93h-1v2.843h1V10.07zm0-4.052h-1V9h1V6.018zm0-3.917h-1V4.96h1V2.101z"/>`),
		g.Group(children),
	)
}