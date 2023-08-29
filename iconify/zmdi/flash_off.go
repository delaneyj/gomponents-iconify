package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m27 24l336 336l-27 27l-89-89l-76 131V237h-64v-79L0 51zm293 149l-33 57L107 49V3h213l-85 170h85z"/>`),
		g.Group(children),
	)
}