package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardBoxLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5v14h16V5H4ZM3 3h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm3 4h2v2H6V7Zm0 4h2v2H6v-2Zm0 4h12v2H6v-2Zm5-4h2v2h-2v-2Zm0-4h2v2h-2V7Zm5 0h2v2h-2V7Zm0 4h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}