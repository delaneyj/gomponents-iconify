package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 8a.5.5 0 0 1-.254.435L7.62 9.924A.75.75 0 0 1 6.5 9.27V6.729a.75.75 0 0 1 1.12-.652l2.626 1.488A.5.5 0 0 1 10.5 8ZM8 2a6 6 0 1 0 0 12A6 6 0 0 0 8 2ZM3 8a5 5 0 1 1 10 0A5 5 0 0 1 3 8Z"/>`),
		g.Group(children),
	)
}