package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FanfictionNet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5h27a10 10 0 0 1 10 10v27h0h-27a10 10 0 0 1-10-10v-27h0Zm0 18.5h37"/>`),
		g.Group(children),
	)
}