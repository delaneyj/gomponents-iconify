package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mivideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.93 42.5a4.06 4.06 0 0 1-2-.53a4 4 0 0 1-2-3.46v-29a4 4 0 0 1 6-3.46l25.13 14.49a4 4 0 0 1 0 6.92L15.93 42a4.07 4.07 0 0 1-2 .5Zm4-26.1v15.2L31.08 24ZM31.08 24l7.99-4.61m-29.13-7.6l7.99 4.61m0 15.2v9.21"/>`),
		g.Group(children),
	)
}