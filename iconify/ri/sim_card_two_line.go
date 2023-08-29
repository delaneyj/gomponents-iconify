package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCardTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4v16h12V7.828L14.172 4H6ZM5 2h10l4.707 4.707a1 1 0 0 1 .293.707V21a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1Zm8 8v8h-2v-6H8v-2h5Zm-5 3h2v2H8v-2Zm6 0h2v2h-2v-2Zm0-3h2v2h-2v-2Zm-6 6h2v2H8v-2Zm6 0h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}