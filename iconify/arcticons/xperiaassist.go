package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xperiaassist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5a10.43 10.43 0 0 0-10.43 10.43c0 5 4.53 9 4.53 14.61m.22 9.66l11.58-2.04M24 4.5a10.43 10.43 0 0 1 10.43 10.43c0 5-4.53 9-4.53 14.61v8.06a5.9 5.9 0 0 1-11.8 0v-3.81l11.8-2.08"/>`),
		g.Group(children),
	)
}