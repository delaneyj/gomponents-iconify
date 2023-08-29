package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RhombusSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.892 4.069A1.75 1.75 0 0 1 5.504 3h7.748a1.75 1.75 0 0 1 1.611 2.431l-2.747 6.502a1.75 1.75 0 0 1-1.612 1.069H2.756a1.75 1.75 0 0 1-1.612-2.432l2.748-6.5Z"/>`),
		g.Group(children),
	)
}