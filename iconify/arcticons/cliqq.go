package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cliqq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.759 43.5L24.241 30.982M11.24 30.5c0 7.222 5.779 13 12.52 13c7.221 0 13-5.778 13-13v-13c0-7.222-5.779-13-13-13s-12.52 5.778-12.52 13v13Z"/><circle cx="24" cy="24" r="3.225" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}