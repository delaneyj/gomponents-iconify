package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditSelectFrameCursorFrameSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5.5h2a1 1 0 0 1 1 1v2m-13 0v-2a1 1 0 0 1 1-1h2m7 13h2a1 1 0 0 0 1-1v-2m-13 0v2a1 1 0 0 0 1 1h2"/>`),
		g.Group(children),
	)
}