package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 320h85v43H0v-43h85q-17 0-29.5-12.5T43 277V64q0-18 12.5-30.5T85 21h342q17 0 29.5 12.5T469 64v213q0 18-12.5 30.5T427 320zM85 64v213h342V64H85z"/>`),
		g.Group(children),
	)
}