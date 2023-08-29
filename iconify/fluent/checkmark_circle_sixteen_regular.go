package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkCircleSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a6 6 0 1 1 0 12A6 6 0 0 1 8 2Zm0 1a5 5 0 1 0 0 10A5 5 0 0 0 8 3Zm-.75 6.042l2.87-2.878a.5.5 0 0 1 .766.637l-.058.07l-3.224 3.232a.5.5 0 0 1-.638.059l-.07-.058l-1.75-1.75a.5.5 0 0 1 .638-.765l.07.057L7.25 9.042l2.87-2.878l-2.87 2.878Z"/>`),
		g.Group(children),
	)
}