package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 38.32L28.609 21L11 3.68L13.72 1L34 21.01L13.72 41z"/>`),
		g.Group(children),
	)
}