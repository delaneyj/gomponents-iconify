package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0zm-6 0a2 2 0 1 0-4 0a2 2 0 0 0 4 0zM4 8a4 4 0 0 1 4-4a.5.5 0 0 0 0-1a5 5 0 0 0-5 5a.5.5 0 0 0 1 0zm9 0a.5.5 0 1 0-1 0a4 4 0 0 1-4 4a.5.5 0 0 0 0 1a5 5 0 0 0 5-5z"/>`),
		g.Group(children),
	)
}