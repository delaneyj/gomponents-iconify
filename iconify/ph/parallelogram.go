package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parallelogram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M245.43 47.31A15.94 15.94 0 0 0 232 40H88.81a16 16 0 0 0-14.59 9.43l-64.8 144A16 16 0 0 0 24 216h143.19a16 16 0 0 0 14.59-9.43l64.8-144a16 16 0 0 0-1.15-15.26ZM167.19 200H24L88.81 56H232Z"/>`),
		g.Group(children),
	)
}