package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6.53 9.02L4.58 7.07l-.88.89l2.83 2.83l.88-.89l4.78-4.77l-.89-.88l-4.77 4.77z"/><path fill="currentColor" d="M8 .5A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.5A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .5zm0 13.75A6.52 6.52 0 0 1 1.25 8A6.52 6.52 0 0 1 8 1.75A6.52 6.52 0 0 1 14.75 8A6.52 6.52 0 0 1 8 14.25z"/>`),
		g.Group(children),
	)
}