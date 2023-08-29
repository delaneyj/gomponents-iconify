package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Capsule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 4.5a5.12 5.12 0 0 0-7.24 0L4.5 12.26a5.12 5.12 0 1 0 7.24 7.24l7.76-7.76a5.12 5.12 0 0 0 0-7.24Zm-9.18 13.59a3.21 3.21 0 0 1-4.41 0a3.13 3.13 0 0 1 0-4.41l3.18-3.18l4.41 4.41Zm7.77-7.77l-3.18 3.18l-4.41-4.41l3.18-3.18a3.12 3.12 0 0 1 4.41 4.41Z"/>`),
		g.Group(children),
	)
}