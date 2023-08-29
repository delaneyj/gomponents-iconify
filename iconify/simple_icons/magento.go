package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magento(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 24l-4.455-2.572v-12l2.97-1.715v12.001l1.485.902l1.485-.902V7.713l2.971 1.715v12L12 24zM22.391 6v12l-2.969 1.714V7.713L12 3.43L4.574 7.713v12.001L1.609 18V6L12 0l10.391 6z"/>`),
		g.Group(children),
	)
}