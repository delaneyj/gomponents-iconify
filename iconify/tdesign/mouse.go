package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 9a8 8 0 1 1 16 0v6a8 8 0 1 1-16 0V9Zm7-5.917A6.002 6.002 0 0 0 6 9h5V3.083Zm2 0V9h5a6.002 6.002 0 0 0-5-5.917ZM18 11H6v4a6 6 0 0 0 12 0v-4Z"/>`),
		g.Group(children),
	)
}