package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatSignalLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3.223a9.003 9.003 0 0 0-5.605 13.592L3 21l4.185-1.395A9.003 9.003 0 0 0 20.777 14m0-4A9.013 9.013 0 0 0 14 3.223M17 12a5 5 0 0 0-5-5m1 5a1 1 0 0 0-1-1"/>`),
		g.Group(children),
	)
}