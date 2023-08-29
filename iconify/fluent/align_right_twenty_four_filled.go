package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 2.75a.75.75 0 0 0-1.5 0v18.5a.75.75 0 0 0 1.5 0V2.75ZM15.75 4A2.25 2.25 0 0 1 18 6.25v2.5A2.25 2.25 0 0 1 15.75 11H5.25A2.25 2.25 0 0 1 3 8.75v-2.5A2.25 2.25 0 0 1 5.25 4h10.5Zm0 9A2.25 2.25 0 0 1 18 15.25v2.5A2.25 2.25 0 0 1 15.75 20h-7a2.25 2.25 0 0 1-2.25-2.25v-2.5A2.25 2.25 0 0 1 8.75 13h7Z"/>`),
		g.Group(children),
	)
}