package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deviantart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.642 4.794l.23-.43V0h-4.367l-.436.44l-2.058 3.925l-.646.435H.015v5.993h4.04l.36.436l-4.18 7.981l-.24.43V24h4.37l.436-.44l2.07-3.925l.644-.436h7.35v-5.993h-4.05l-.36-.438l4.186-7.977z"/>`),
		g.Group(children),
	)
}