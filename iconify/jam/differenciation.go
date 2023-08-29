package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Differenciation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 .674a7 7 0 1 1 0 12.653A7 7 0 1 1 10 .673ZM10 3a4.992 4.992 0 0 0-2 4c0 1.636.786 3.088 2 4a4.992 4.992 0 0 0 2-4a4.992 4.992 0 0 0-2-4Z"/>`),
		g.Group(children),
	)
}