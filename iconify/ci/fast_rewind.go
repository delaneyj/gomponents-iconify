package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastRewind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 18l-8.5-6L20 6v12Zm-9 0l-8.5-6L11 6v12Z"/>`),
		g.Group(children),
	)
}