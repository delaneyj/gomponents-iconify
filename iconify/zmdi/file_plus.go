package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m213 3l128 128v256q0 17-12.5 29.5T299 429H42q-17 0-29.5-12.5T0 387V45q0-17 12.5-29.5T43 3h170zm43 298v-42h-64v-64h-43v64H85v42h64v64h43v-64h64zm-64-149h117L192 35v117z"/>`),
		g.Group(children),
	)
}