package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteBinThreeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 7v13a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V7H2V5h20v2h-2ZM6 7v13h12V7H6Zm5 2h2v2h-2V9Zm0 3h2v2h-2v-2Zm0 3h2v2h-2v-2ZM7 2h10v2H7V2Z"/>`),
		g.Group(children),
	)
}