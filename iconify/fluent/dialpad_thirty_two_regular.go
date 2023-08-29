package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialpadThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 5a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm7-14a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm7-21a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/>`),
		g.Group(children),
	)
}