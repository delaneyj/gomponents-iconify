package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M25 2.75a.75.75 0 0 0-1.5 0v22.5a.75.75 0 0 0 1.5 0V2.75ZM19.25 5A2.75 2.75 0 0 1 22 7.75v2.5A2.75 2.75 0 0 1 19.25 13H5.75A2.75 2.75 0 0 1 3 10.25v-2.5A2.75 2.75 0 0 1 5.75 5h13.5Zm0 10A2.75 2.75 0 0 1 22 17.75v2.5A2.75 2.75 0 0 1 19.25 23h-8a2.75 2.75 0 0 1-2.75-2.75v-2.5A2.75 2.75 0 0 1 11.25 15h8Z"/>`),
		g.Group(children),
	)
}