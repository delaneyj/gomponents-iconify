package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 2.75a.75.75 0 0 1 1.5 0v22.5a.75.75 0 0 1-1.5 0V2.75ZM8.75 5A2.75 2.75 0 0 0 6 7.75v2.5A2.75 2.75 0 0 0 8.75 13h13.5A2.75 2.75 0 0 0 25 10.25v-2.5A2.75 2.75 0 0 0 22.25 5H8.75Zm0 10A2.75 2.75 0 0 0 6 17.75v2.5A2.75 2.75 0 0 0 8.75 23h8a2.75 2.75 0 0 0 2.75-2.75v-2.5A2.75 2.75 0 0 0 16.75 15h-8Z"/>`),
		g.Group(children),
	)
}