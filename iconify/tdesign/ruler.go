package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.719 2h11.035L13.64 20H22v2H1.719l5-20Zm4.869 18l3.657-16H8.28l-.625 2.5h4.22v2h-4.72L6.53 11h4.22v2H6.03l-.625 2.5h4.22v2h-4.72L4.28 20h7.308Z"/>`),
		g.Group(children),
	)
}