package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M43.634 4.366a1.25 1.25 0 0 1 0 1.768l-37.5 37.5a1.25 1.25 0 0 1-1.768-1.768l37.5-37.5a1.25 1.25 0 0 1 1.768 0Z"/>`),
		g.Group(children),
	)
}