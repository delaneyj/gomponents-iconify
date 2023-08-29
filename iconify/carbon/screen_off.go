package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 22H11.41L30 3.41L28.59 2l-2 2H4a2 2 0 0 0-2 2v16h2V6h20.59L2 28.59L3.41 30l6-6H12v4H8v2h16v-2h-4v-4h8a2 2 0 0 0 2-2V9h-2Zm-10 6h-4v-4h4Z"/>`),
		g.Group(children),
	)
}