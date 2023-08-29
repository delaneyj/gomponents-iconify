package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareTopUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m18.432 8.75l3.205-3.205C22 7.009 22 9.04 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12s0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c2.96 0 4.991 0 6.455.363L15.25 5.568V5a2.25 2.25 0 0 0-4.5 0v6A2.25 2.25 0 0 0 13 13.25h6a2.25 2.25 0 0 0 0-4.5h-.568Z"/><path d="M19 11.75a.75.75 0 0 0 0-1.5h-4.19l6.72-6.72a.75.75 0 0 0-1.06-1.06l-6.72 6.72V5a.75.75 0 0 0-1.5 0v6c0 .414.336.75.75.75h6Z"/></g>`),
		g.Group(children),
	)
}