package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitedSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 3.5a4.5 4.5 0 0 0-3.668 7.107l6.275-6.275A4.48 4.48 0 0 0 8 3.5Zm3.668 1.893l-6.275 6.275a4.5 4.5 0 0 0 6.276-6.276ZM2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8Z"/>`),
		g.Group(children),
	)
}