package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudShield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.42 7.22A7 7 0 0 0 5.06 9.11a4 4 0 0 0-.38 7.66a1.13 1.13 0 0 0 .32.05a1 1 0 0 0 .32-2A2 2 0 0 1 4 13a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 12a5 5 0 0 0-3.58-4.78Zm-3.63 4.44a2.73 2.73 0 0 1-2.2-.47a1 1 0 0 0-1.18 0a2.72 2.72 0 0 1-2.2.47a1 1 0 0 0-.84.2a1 1 0 0 0-.37.77V16a4.63 4.63 0 0 0 1.84 3.7l1.57 1.15a1 1 0 0 0 1.18 0l1.57-1.16A4.6 4.6 0 0 0 16 16v-3.37a1 1 0 0 0-.37-.77a1 1 0 0 0-.84-.2ZM14 16a2.62 2.62 0 0 1-1 2l-1 .72l-1-.72a2.62 2.62 0 0 1-1-2v-2.28a4.68 4.68 0 0 0 2-.55a4.68 4.68 0 0 0 2 .55Z"/>`),
		g.Group(children),
	)
}