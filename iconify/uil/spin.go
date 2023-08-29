package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 3a7 7 0 0 0 0 14a5 5 0 0 0 0-10a3 3 0 0 0 0 6a1 1 0 0 0 0-2a1 1 0 0 1 0-2a3 3 0 0 1 0 6a5 5 0 0 1 0-10a7 7 0 0 1 0 14a9 9 0 0 1-9-9a1 1 0 0 0-2 0a11 11 0 0 0 11 11a9 9 0 0 0 0-18Z"/>`),
		g.Group(children),
	)
}