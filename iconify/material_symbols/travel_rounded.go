package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.85 17.15l-2.5-1.375q-.275-.15-.312-.437t.162-.488l.3-.3q.1-.1.238-.138T5 14.4l2.2.3l3.9-3.9l-6.8-3.7q-.375-.2-.437-.625T4.1 5.75l.25-.25q.175-.175.388-.225t.437 0L14.25 7.6l3.925-3.875Q18.6 3.3 19.238 3.3t1.062.425q.425.425.425 1.063T20.3 5.85l-3.9 3.9l2.325 9.025q.05.25-.013.488t-.237.412l-.125.125q-.35.35-.813.275t-.687-.5L13.2 12.9l-3.9 3.9l.3 2.15q.025.175-.025.325t-.175.275l-.125.125q-.25.25-.6.212t-.525-.362l-1.3-2.375Z"/>`),
		g.Group(children),
	)
}