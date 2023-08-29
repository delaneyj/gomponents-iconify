package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReorderTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.5 8a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1h-15Zm0 3a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1h-15Z"/>`),
		g.Group(children),
	)
}