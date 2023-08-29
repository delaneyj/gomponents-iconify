package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatClearRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.2 10.35l-2.325-2.325L7.85 5H18.5q.625 0 1.063.438T20 6.5q0 .625-.438 1.063T18.5 8h-4.3l-1 2.35Zm5.9 11.55l-7.6-7.6l-1.6 3.775q-.175.425-.563.675T8.5 19q-.8 0-1.25-.675T7.125 16.9L9.2 12L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}