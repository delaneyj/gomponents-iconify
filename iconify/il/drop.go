package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M515 206q45 45 69 103t26 118t-21 117t-67 103t-101 67t-116 23t-115-23t-102-67t-67-103T0 427t25-118t69-103L296 4q8-8 17 0z"/>`),
		g.Group(children),
	)
}