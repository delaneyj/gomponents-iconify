package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveHA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.783 12.208a2.608 2.608 0 1 1 5.215 0a2.608 2.608 0 0 1-5.215.001zm-9.392 0a2.609 2.609 0 1 1 5.217 0a2.609 2.609 0 0 1-5.217.001zm-9.391 0a2.609 2.609 0 1 1 2.609 2.61h-.001A2.608 2.608 0 0 1 0 12.21v-.002z"/>`),
		g.Group(children),
	)
}