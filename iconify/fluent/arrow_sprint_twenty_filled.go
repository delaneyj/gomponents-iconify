package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSprintTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 6.5a3 3 0 1 0 0 6h6.44l-.72-.72a.75.75 0 1 1 1.06-1.06l2 2a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 1 1-1.06-1.06l.72-.72H10a4.5 4.5 0 1 1 4.032-2.5h-1.796A3 3 0 0 0 10 6.5Zm-7.25 6h2.64A5.53 5.53 0 0 0 6.836 14H2.75a.75.75 0 0 1 0-1.5Z"/>`),
		g.Group(children),
	)
}