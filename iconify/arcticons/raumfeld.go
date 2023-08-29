package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Raumfeld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.992 40.465c3.188.948 13.431 2.585 14.858-8.788s2.805-24.815 2.805-24.815"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 6.862H14.325c-8.078 0-12.02 4.093-5.687 13.14"/>`),
		g.Group(children),
	)
}