package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineExpand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20h16v2H4zM4 2h16v2H4zm9 7h3l-4-4l-4 4h3v6H8l4 4l4-4h-3z"/>`),
		g.Group(children),
	)
}