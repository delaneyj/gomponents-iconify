package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlashForwardTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.658 1.026a.5.5 0 0 1 .316.632l-3 9a.5.5 0 1 1-.948-.316l3-9a.5.5 0 0 1 .632-.316Z"/>`),
		g.Group(children),
	)
}