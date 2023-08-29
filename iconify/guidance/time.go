package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Time(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3 .5h18m-18 23h18m-15.5 0v-6l2.856-1.714a4.415 4.415 0 0 0 0-7.572L5.5 6.5v-6m13 0v6l-2.856 1.714a4.416 4.416 0 0 0 0 7.572L18.5 17.5v6"/>`),
		g.Group(children),
	)
}