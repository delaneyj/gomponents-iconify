package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReflectVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3 18l6-6l-6-6v12zm12-6l6 6V6l-6 6zm-4-9h2v3h-2zm0 5h2v3h-2zm0 5h2v3h-2zm0 5h2v3h-2z"/>`),
		g.Group(children),
	)
}