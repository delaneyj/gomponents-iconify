package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioDescription(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m2.75 16.25l1.605-2.841m0 0L7.408 8H8.75v5.409m-4.395 0H4.25m.105 0H8.75m0 0V16.5m9.25-9s.5 1.5.5 4.5s-.5 4.5-.5 4.5m2.5-9S21 9 21 12s-.5 4.5-.5 4.5m-19 4S.5 17 .5 12s1-8.5 1-8.5h21s1 3.5 1 8.5s-1 8.5-1 8.5h-21ZM11 16V8h3.75c.644 0 1.25.522 1.25 1.167v5.666c0 .645-.606 1.167-1.25 1.167H11Z"/>`),
		g.Group(children),
	)
}