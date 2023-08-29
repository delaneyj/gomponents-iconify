package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="m25.496 10.088l-2.622-2.622V3h2.25v3.534l1.964 1.964l-1.592 1.59z"/><path fill="currentColor" d="M24 1a6 6 0 1 0 6 6a6.007 6.007 0 0 0-6-6Zm1.497 9.088l-2.622-2.622V3h2.25v3.534l1.964 1.964Z"/><path fill="currentColor" d="M6 16v-6h9V8H6.184A2.995 2.995 0 0 1 9 6h6V4H9a5.006 5.006 0 0 0-5 5v12a4.99 4.99 0 0 0 3.582 4.77L5.769 30h2.176l1.714-4h8.682l1.714 4h2.176l-1.813-4.23A4.99 4.99 0 0 0 24 21v-5Zm16 4h-3v2h2.816A2.995 2.995 0 0 1 19 24H9a2.995 2.995 0 0 1-2.816-2H9v-2H6v-2h16Z"/>`),
		g.Group(children),
	)
}