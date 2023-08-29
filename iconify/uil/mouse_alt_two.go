package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a7 7 0 0 0-7 7v6a7 7 0 0 0 14 0V9a7 7 0 0 0-7-7Zm5 13a5 5 0 0 1-10 0V9a5 5 0 0 1 4-4.9V12a1 1 0 0 0 2 0V4.1A5 5 0 0 1 17 9Z"/>`),
		g.Group(children),
	)
}