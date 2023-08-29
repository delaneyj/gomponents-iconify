package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.583 33.75H17.5v5.417a1.083 1.083 0 0 1-1.083 1.083H5.583A1.083 1.083 0 0 1 4.5 39.167v-4.334a1.083 1.083 0 0 1 1.083-1.083ZM24 20.75v11.917a1.083 1.083 0 0 1-1.083 1.083H17.5h0V21.833a1.083 1.083 0 0 1 1.083-1.083H24Zm18.417-6.5H30.5h0V8.833a1.083 1.083 0 0 1 1.083-1.083h10.834A1.083 1.083 0 0 1 43.5 8.833v4.334a1.083 1.083 0 0 1-1.083 1.083ZM24 27.25V15.333a1.083 1.083 0 0 1 1.083-1.083H30.5v11.917a1.083 1.083 0 0 1-1.083 1.083H24h0Z"/>`),
		g.Group(children),
	)
}