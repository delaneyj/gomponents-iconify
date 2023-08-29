package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlashForwardSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.44 2.037a.5.5 0 0 1 .273.652l-4.5 11a.5.5 0 0 1-.926-.378l4.5-11a.5.5 0 0 1 .652-.274Z"/>`),
		g.Group(children),
	)
}