package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zarchiver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.89 7.05h23.53l-11 25.42h13.17L34.26 41H4.51l11-25.42h-5L4.89 7.05Zm23.53 0L43.51 41h-9.25L24 17.25"/>`),
		g.Group(children),
	)
}