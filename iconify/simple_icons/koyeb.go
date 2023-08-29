package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Koyeb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 12.822V7.044L11.985.115L24 7.037v5.77L11.992 5.892Zm11.985 1.114L1.92 19.759L0 18.645v-3.557l11.985-6.93L24 15.089v3.542l-1.92 1.13Zm-3.028 9.949L3.95 21.004l8.036-4.656l8.066 4.656l-5.009 2.88l-3.05-1.759Z"/>`),
		g.Group(children),
	)
}