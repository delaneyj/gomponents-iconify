package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldLessRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 16.9l-2.4 2.4q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l3.1-3.1q.15-.15.325-.212t.375-.063q.2 0 .375.063t.325.212l3.1 3.1q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L12 16.9Zm0-9.8l2.4-2.4q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-3.1 3.1q-.15.15-.325.213T12 9.475q-.2 0-.375-.062T11.3 9.2L8.2 6.1q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275L12 7.1Z"/>`),
		g.Group(children),
	)
}