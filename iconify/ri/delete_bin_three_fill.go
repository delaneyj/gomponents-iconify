package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteBinThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 7v14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V7H2V5h20v2h-2Zm-9 2v2h2V9h-2Zm0 3v2h2v-2h-2Zm0 3v2h2v-2h-2ZM7 2h10v2H7V2Z"/>`),
		g.Group(children),
	)
}