package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageFrame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 1h-2v2H9v2H7v2H2v16h20V7h-5V5h-2V3h-2V1zm2 6H9V5h2V3h2v2h2v2zM4 9h16v12H4V9zm10 6v-2h-2v2h-2v2H8v2h2v-2h2v-2h2zm2 2v-2h-2v2h2zm0 0v2h2v-2h-2zM6 13v-2h2v2H6z"/>`),
		g.Group(children),
	)
}