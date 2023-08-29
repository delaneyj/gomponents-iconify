package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderMale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.5 3H21v7.5h-2V6.414l-3.394 3.394a7 7 0 1 1-1.414-1.414L17.586 5H13.5V3ZM10 9a5 5 0 1 0 0 10a5 5 0 0 0 0-10Z"/>`),
		g.Group(children),
	)
}