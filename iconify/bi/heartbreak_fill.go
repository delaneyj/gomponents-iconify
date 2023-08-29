package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartbreakFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.931.586L7 3l1.5 4l-2 3L8 15C22.534 5.396 13.757-2.21 8.931.586ZM7.358.77L5.5 3L7 7l-1.5 3l1.815 4.537C-6.533 4.96 2.685-2.467 7.358.77Z"/>`),
		g.Group(children),
	)
}