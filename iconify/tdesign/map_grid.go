package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h16V4H4Zm5 2v1h2V6h2v1h2V6h2v1h1v2h-1v2h1v2h-1v2h1v2h-1v1h-2v-1h-2v1h-2v-1H9v1H7v-1H6v-2h1v-2H6v-2h1V9H6V7h1V6h2Zm0 3v2h2V9H9Zm4 0v2h2V9h-2Zm2 4h-2v2h2v-2Zm-4 2v-2H9v2h2Z"/>`),
		g.Group(children),
	)
}