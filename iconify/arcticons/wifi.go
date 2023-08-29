package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.765 28.616a10.515 10.515 0 0 0-13.435-.037v.037m6.691 4.305a4.216 4.216 0 1 0 4.216 4.216h0a4.216 4.216 0 0 0-4.216-4.216Zm13.157-11.595a20.675 20.675 0 0 0-26.314 0M43.5 13.674a30.568 30.568 0 0 0-39 0"/>`),
		g.Group(children),
	)
}