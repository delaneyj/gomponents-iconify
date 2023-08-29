package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 7a5 5 0 0 0 0 10h10a5 5 0 0 0 0-10H7Zm0 2.5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}