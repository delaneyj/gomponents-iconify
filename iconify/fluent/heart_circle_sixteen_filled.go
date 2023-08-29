package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartCircleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8Zm6-1l-.422-.492a1.465 1.465 0 1 0-2.156 1.98l2.4 2.44c.097.1.258.1.356 0l2.4-2.44a1.465 1.465 0 1 0-2.157-1.98L8 7Z"/>`),
		g.Group(children),
	)
}