package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 3.748a.75.75 0 0 1 .75-.75h16.504a.75.75 0 0 1 .6 1.2L16.69 9.75l4.164 5.551a.75.75 0 0 1-.6 1.2H4.5v4.75a.75.75 0 0 1-.648.743L3.75 22a.75.75 0 0 1-.743-.648L3 21.25V3.748Zm15.754.75H4.5V15h14.254l-3.602-4.8a.75.75 0 0 1 0-.9l3.602-4.802Z"/>`),
		g.Group(children),
	)
}