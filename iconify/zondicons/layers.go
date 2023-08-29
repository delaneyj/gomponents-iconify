package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10 1l10 6l-10 6L0 7l10-6zm6.67 10L20 13l-10 6l-10-6l3.33-2L10 15l6.67-4z"/>`),
		g.Group(children),
	)
}