package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 64q17 0 29.5 12.5T469 107v170q0 18-12.5 30.5T427 320H43q-18 0-30.5-12.5T0 277V107q0-18 12.5-30.5T43 64h384zm0 213V107h-43v85h-43v-85h-42v85h-43v-85h-43v85h-42v-85h-43v85H85v-85H43v170h384z"/>`),
		g.Group(children),
	)
}