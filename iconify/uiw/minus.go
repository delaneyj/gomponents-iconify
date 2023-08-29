package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M.7 10.4A.7.7 0 0 1 .7 9h18.606a.7.7 0 0 1 0 1.4H.7Z"/>`),
		g.Group(children),
	)
}