package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DarkThemeTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 3a7 7 0 1 1 0 14V3Zm0-1a8 8 0 1 0 0 16a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}