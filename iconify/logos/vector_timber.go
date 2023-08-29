package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorTimber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#10E7FF" d="m256 0l-91.429 150.862l-18.969-31.219L217.61 0H256ZM47.204 0l44.674 73.367L136.53 0h45.365L90.947 150.068L0 0h47.204Z"/>`),
		g.Group(children),
	)
}