package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flamingo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.978 14.75v18.5L24 42.5l16.021-9.25v-18.5L24 5.5L7.978 14.75z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 42.5V24l16.021 9.25m0-18.5L24 24V5.5M7.978 14.75L24 24L7.978 33.25"/>`),
		g.Group(children),
	)
}