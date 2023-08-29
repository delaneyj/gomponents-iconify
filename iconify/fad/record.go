package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Record(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M128 209c-44.735 0-81-36.265-81-81s36.265-81 81-81s81 36.265 81 81s-36.265 81-81 81zm.5-33c26.51 0 48-21.49 48-48s-21.49-48-48-48s-48 21.49-48 48s21.49 48 48 48z"/>`),
		g.Group(children),
	)
}