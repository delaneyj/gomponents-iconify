package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.28 3.22a.75.75 0 0 0-1.06 1.06L14.44 15.5H6.75a.75.75 0 0 0 0 1.5h9.5a.747.747 0 0 0 .75-.75v-9.5a.75.75 0 0 0-1.5 0v7.69L4.28 3.22Z"/>`),
		g.Group(children),
	)
}