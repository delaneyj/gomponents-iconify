package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backspace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.6 2.5H4.49a1.25 1.25 0 0 0-1 .51L.39 7.26a1.26 1.26 0 0 0 0 1.48L3.48 13a1.26 1.26 0 0 0 1 .51H14.6a1.25 1.25 0 0 0 1.25-1.25V3.75A1.25 1.25 0 0 0 14.6 2.5zm0 9.75H4.49L1.4 8l3.09-4.25H14.6z"/><path fill="currentColor" d="m7.86 10.55l1.99-1.72l1.99 1.72l.82-.94L10.81 8l1.85-1.61l-.82-.94l-1.99 1.72l-1.99-1.72l-.82.94L8.9 8L7.04 9.61l.82.94z"/>`),
		g.Group(children),
	)
}