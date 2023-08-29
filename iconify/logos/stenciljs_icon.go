package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StenciljsIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="m193.065 138.495l-53.333 58.428H41.599L94.4 138.495h98.664ZM256 69.248l-53.305 58.427H0l53.305-58.427H256ZM214.399 0l-53.087 58.428h-98.38L116.1 0h98.3Z"/>`),
		g.Group(children),
	)
}