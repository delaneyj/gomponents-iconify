package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OptionsSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 9a2 2 0 0 1 1.937 1.5H13.5a.5.5 0 0 1 .09.992l-.09.008l-5.563.001a2 2 0 0 1-3.874 0L2.5 11.5a.5.5 0 0 1-.09-.992l.09-.008h1.563A2 2 0 0 1 6 9Zm4-6a2 2 0 0 1 1.937 1.5H13.5a.5.5 0 0 1 .09.992l-.09.008l-1.563.001a2 2 0 0 1-3.874 0L2.5 5.5a.5.5 0 0 1-.09-.992L2.5 4.5h5.563A2 2 0 0 1 10 3Z"/>`),
		g.Group(children),
	)
}