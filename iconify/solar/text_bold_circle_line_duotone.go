package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBoldCircleLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Z" opacity=".5"/><path d="M8 7.522C8 6.682 8.681 6 9.522 6H12a3 3 0 1 1 0 6H8V7.522ZM8 12h5a3 3 0 1 1 0 6H9.176C8.526 18 8 17.473 8 16.823V12Z"/></g>`),
		g.Group(children),
	)
}