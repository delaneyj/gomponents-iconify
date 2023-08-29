package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanCableSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 6.5V5h1v1.5a.5.5 0 0 1-1 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M9 0h2v4h1v3.618L9.809 12H9v3H8v-3H7v3H6v-3h-.809L3 7.618V4h1V0h2v3h3V0Zm0 4H6v2.5a1.5 1.5 0 1 0 3 0V4ZM6 9v1h3V9H6Z" clip-rule="evenodd"/><path fill="currentColor" d="M8 0H7v2h1V0Z"/>`),
		g.Group(children),
	)
}