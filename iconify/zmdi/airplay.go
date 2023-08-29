package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m107 405l128-128l128 128H107zM427 0q17 0 29.5 12.5T469 43v256q0 17-12.5 29.5T427 341h-86v-42h86V43H43v256h85v42H43q-18 0-30.5-12.5T0 299V43q0-18 12.5-30.5T43 0h384z"/>`),
		g.Group(children),
	)
}