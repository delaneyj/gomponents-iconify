package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeafOneSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.275 4.536A4.886 4.886 0 0 1 8.5 13.022V14.5a.5.5 0 1 1-1 0v-1.478a4.886 4.886 0 0 1-2.775-8.486l2.437-2.2a1.25 1.25 0 0 1 1.676 0l2.437 2.2ZM8.5 7.501a.5.5 0 0 0-1 0v4.514a3.91 3.91 0 0 0 1 0V7.501Z"/>`),
		g.Group(children),
	)
}