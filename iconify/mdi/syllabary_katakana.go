package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyllabaryKatakana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4v2h14v.09l-3.71 3.7l1.42 1.42L20 6.91V4m-9 5v4c0 2.78-.75 3.89-2.64 5.46L9.64 20c2.11-1.76 3.36-3.78 3.36-7V9Z"/>`),
		g.Group(children),
	)
}