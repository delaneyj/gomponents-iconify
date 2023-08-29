package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineExpand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20h16v2H4zM4 2h16v2H4zm5.41 11.59L8 15l4 4l4-4l-1.41-1.41L13 15.17V8.83l1.59 1.58L16 9l-4-4l-4 4l1.41 1.41L11 8.83v6.34z"/>`),
		g.Group(children),
	)
}