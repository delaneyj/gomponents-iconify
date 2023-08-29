package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Obimy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.1 31.75c-1.6 1.5-3.7 2.4-5.9 2.4s-4.4-.9-6-2.4m3.9-13.7c-1.4-1.4-3.3-2.2-5.3-2.2c-1.9 0-3.8.8-5.2 2.1m24.8.1c-1.4-1.4-3.3-2.2-5.3-2.2c-1.9 0-3.8.8-5.2 2.1"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="9" ry="9"/>`),
		g.Group(children),
	)
}