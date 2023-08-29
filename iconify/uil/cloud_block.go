package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudBlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.42 7.72A7 7 0 0 0 5.06 9.61a4 4 0 0 0-.38 7.66a1.13 1.13 0 0 0 .32.05a1 1 0 0 0 .32-2A2 2 0 0 1 4 13.5a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 12.5a5 5 0 0 0-3.58-4.78Zm-9.25 6a4 4 0 1 0 5.66 0a4.1 4.1 0 0 0-5.66-.05ZM10 16.5a2 2 0 0 1 2-2a2.09 2.09 0 0 1 .51.07L10.07 17a2.09 2.09 0 0 1-.07-.5Zm3.41 1.41a2 2 0 0 1-1.91.5L13.93 16a2.09 2.09 0 0 1 .07.51a2 2 0 0 1-.59 1.4Z"/>`),
		g.Group(children),
	)
}