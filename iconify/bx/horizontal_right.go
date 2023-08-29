package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 11H5v2h8v3l4-4l-4-4v3zm6-8h2v3h-2zm0 5h2v3h-2zm0 5h2v3h-2zm0 5h2v3h-2z"/>`),
		g.Group(children),
	)
}