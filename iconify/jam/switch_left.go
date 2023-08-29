package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2a5 5 0 1 0 0 10h6a5 5 0 0 0 0-10H7zm0-2h6a7 7 0 0 1 0 14H7A7 7 0 0 1 7 0zm0 11a4 4 0 1 1 0-8a4 4 0 0 1 0 8zm0-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4z"/>`),
		g.Group(children),
	)
}