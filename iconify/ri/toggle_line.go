package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 7a5 5 0 0 0 0 10h8a5 5 0 0 0 0-10H8Zm0-2h8a7 7 0 1 1 0 14H8A7 7 0 1 1 8 5Zm0 10a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		g.Group(children),
	)
}