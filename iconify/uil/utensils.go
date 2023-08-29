package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Utensils(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2a1 1 0 0 0-1 1v5.46l-1 .67V3a1 1 0 0 0-2 0v6.13l-1-.67V3a1 1 0 0 0-2 0v6a1 1 0 0 0 .45.83L15 11.54V21a1 1 0 0 0 2 0v-9.46l2.55-1.71A1 1 0 0 0 20 9V3a1 1 0 0 0-1-1ZM9 2a5 5 0 0 0-5 5v6a1 1 0 0 0 1 1h3v7a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1ZM8 12H6V7a3 3 0 0 1 2-2.83Z"/>`),
		g.Group(children),
	)
}