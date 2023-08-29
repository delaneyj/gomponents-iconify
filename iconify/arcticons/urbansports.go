package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Urbansports(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25 4.5H10.8v25.8c0 7.3 5.9 13.2 13.2 13.2c7.3 0 13.2-5.9 13.2-13.2V17.7m0-5.3V4.5h-7.6"/>`),
		g.Group(children),
	)
}