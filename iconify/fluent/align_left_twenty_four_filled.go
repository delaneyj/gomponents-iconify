package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 2.75a.75.75 0 0 1 1.5 0v18.5a.75.75 0 0 1-1.5 0V2.75ZM8.25 4A2.25 2.25 0 0 0 6 6.25v2.5A2.25 2.25 0 0 0 8.25 11h10.5A2.25 2.25 0 0 0 21 8.75v-2.5A2.25 2.25 0 0 0 18.75 4H8.25Zm0 9A2.25 2.25 0 0 0 6 15.25v2.5A2.25 2.25 0 0 0 8.25 20h7a2.25 2.25 0 0 0 2.25-2.25v-2.5A2.25 2.25 0 0 0 15.25 13h-7Z"/>`),
		g.Group(children),
	)
}