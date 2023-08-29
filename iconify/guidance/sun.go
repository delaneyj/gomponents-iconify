package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><path d="M12 24a4.802 4.802 0 0 1 3.11-4.494a4.802 4.802 0 0 1 5.375.98a4.802 4.802 0 0 1-.979-5.377A4.802 4.802 0 0 1 24 12a4.802 4.802 0 0 1-4.494-3.11a4.802 4.802 0 0 1 .98-5.375a4.802 4.802 0 0 1-5.377.979A4.802 4.802 0 0 1 12 0a4.802 4.802 0 0 1-3.11 4.494a4.802 4.802 0 0 1-5.375-.98a4.802 4.802 0 0 1 .979 5.377A4.802 4.802 0 0 1 0 12a4.802 4.802 0 0 1 4.494 3.11a4.802 4.802 0 0 1-.98 5.375a4.802 4.802 0 0 1 5.377-.979A4.802 4.802 0 0 1 12 24Z"/><path d="M16 12a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/></g>`),
		g.Group(children),
	)
}