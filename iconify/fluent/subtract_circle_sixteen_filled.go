package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractCircleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a6 6 0 1 0 0 12A6 6 0 0 0 8 2ZM5.5 7.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}