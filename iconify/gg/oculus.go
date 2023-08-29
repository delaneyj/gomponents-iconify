package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oculus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 10H8a2 2 0 1 0 0 4h8a2 2 0 1 0 0-4ZM8 6a6 6 0 1 0 0 12h8a6 6 0 0 0 0-12H8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}