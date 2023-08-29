package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scrimba(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 6.222a2.222 2.222 0 0 1-2.222 2.222h-8.89a2.222 2.222 0 0 1 0-4.444h8.89C23.005 4 24 4.995 24 6.222zm-7.333 9.334h-8.89a2.222 2.222 0 0 0 0 4.444h8.89a2.222 2.222 0 0 0 0-4.444zm0-5.778H13.11a2.222 2.222 0 0 0 0 4.444h3.556a2.222 2.222 0 0 0 0-4.444zM2.222 15.556a2.222 2.222 0 1 0 0 4.444a2.222 2.222 0 0 0 0-4.444z"/>`),
		g.Group(children),
	)
}