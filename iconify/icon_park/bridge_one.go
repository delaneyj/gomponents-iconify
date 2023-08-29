package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BridgeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 19V33H12C12 26.3726 17.3726 21 24 21C30.6274 21 36 26.3726 36 33H44V19C44 19 31.9652 15 24 15C16.0348 15 4 19 4 19Z"/>`),
		g.Group(children),
	)
}