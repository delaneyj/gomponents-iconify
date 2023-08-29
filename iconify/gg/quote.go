package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.135 9h3L10 14.607H7L9.135 9Zm5 0h3L15 14.607h-3L14.135 9Z"/>`),
		g.Group(children),
	)
}