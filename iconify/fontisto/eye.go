package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 0C7.164 0 0 11.844 0 11.844S7.164 24 16 24s16-12.156 16-12.156S24.836 0 16 0zm0 20a8 8 0 1 1 0-16a8 8 0 0 1 0 16z"/><path fill="currentColor" d="M20 12.016a4 4 0 1 1-8 0a4 4 0 0 1 8 0z"/>`),
		g.Group(children),
	)
}