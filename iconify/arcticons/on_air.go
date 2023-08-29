package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OnAir(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 24c0 11.874-9.626 21.5-21.5 21.5S2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24Zm-33.477-7.032h10.212m-5.106 15.709v-15.71"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.977 16.968l-5.106 15.71l-5.106-15.71"/>`),
		g.Group(children),
	)
}