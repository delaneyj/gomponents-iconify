package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 20.5h-4a.76.76 0 0 1-.75-.75V12L3.9 4.69a.73.73 0 0 1-.07-.78a.76.76 0 0 1 .67-.41h15a.76.76 0 0 1 .67.41a.73.73 0 0 1-.07.78L14.75 12v7.75a.76.76 0 0 1-.75.75ZM10.75 19h2.5v-7.25a.71.71 0 0 1 .15-.44L18 5H6l4.62 6.31a.71.71 0 0 1 .15.44Z"/>`),
		g.Group(children),
	)
}