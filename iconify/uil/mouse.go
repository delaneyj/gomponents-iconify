package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a7 7 0 0 0-7 7v6a7 7 0 0 0 14 0V9a7 7 0 0 0-7-7ZM7 9a5 5 0 0 1 4-4.9V10H7Zm10 6a5 5 0 0 1-10 0v-3h10Zm0-5h-4V4.1A5 5 0 0 1 17 9Z"/>`),
		g.Group(children),
	)
}