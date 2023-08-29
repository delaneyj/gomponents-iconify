package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qfield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.063 22.575l-2.586 10.517l-10.683 2.503l19.85 6.637l-6.58-19.657h-.001Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.692 40.577A21.412 21.412 0 0 1 24 45.5C12.126 45.5 2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24c0 4.981-1.694 9.566-4.537 13.211"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.665 37.893A14.957 14.957 0 0 1 24 39c-8.284 0-15-6.716-15-15S15.716 9 24 9s15 6.716 15 15c0 1.716-.288 3.364-.819 4.9"/>`),
		g.Group(children),
	)
}