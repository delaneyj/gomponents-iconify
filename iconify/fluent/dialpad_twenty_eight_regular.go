package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialpadTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 4.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 6a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 6a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm6.5-12a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 6a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 6a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 6a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm6.5-18a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 6a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 6a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/>`),
		g.Group(children),
	)
}