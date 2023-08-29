package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloverThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 2a6.5 6.5 0 0 0 0 13H14a1 1 0 0 0 1-1V8.5A6.5 6.5 0 0 0 8.5 2ZM4 8.5a4.5 4.5 0 0 1 9 0V13H8.5A4.5 4.5 0 0 1 4 8.5ZM8.5 30a6.5 6.5 0 1 1 0-13H14a1 1 0 0 1 1 1v5.5A6.5 6.5 0 0 1 8.5 30ZM4 23.5a4.5 4.5 0 1 0 9 0V19H8.5A4.5 4.5 0 0 0 4 23.5Zm26-15a6.5 6.5 0 1 0-13 0V14a1 1 0 0 0 1 1h5.5A6.5 6.5 0 0 0 30 8.5ZM23.5 4a4.5 4.5 0 1 1 0 9H19V8.5A4.5 4.5 0 0 1 23.5 4Zm0 26a6.5 6.5 0 1 0 0-13H18a1 1 0 0 0-1 1v5.5a6.5 6.5 0 0 0 6.5 6.5Zm4.5-6.5a4.5 4.5 0 1 1-9 0V19h4.5a4.5 4.5 0 0 1 4.5 4.5Z"/>`),
		g.Group(children),
	)
}