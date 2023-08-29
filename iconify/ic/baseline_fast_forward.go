package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineFastForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4 18l8.5-6L4 6v12zm9-12v12l8.5-6L13 6z"/>`),
		g.Group(children),
	)
}