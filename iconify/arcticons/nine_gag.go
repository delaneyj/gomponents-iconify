package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NineGag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.93 14.25L24 24L7.07 14.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.07 14.25L24 4.5l16.93 9.75v19.5L24 43.5L7.07 33.75"/>`),
		g.Group(children),
	)
}