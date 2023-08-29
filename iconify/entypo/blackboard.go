package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blackboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.539 20H6l1.406-3.698l-2.966-1.004L2.539 20zm10.055-3.698L14 20h3.461l-1.901-4.702l-2.966 1.004zM18 2h-6.5L11 0H9l-.5 2H2a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}