package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kasa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.369 6.654L6.85 22.17a4.614 4.614 0 0 0-1.35 3.263v14.282a2.307 2.307 0 0 0 2.307 2.307h32.386a2.307 2.307 0 0 0 2.307-2.307V25.433a4.614 4.614 0 0 0-1.351-3.262L25.63 6.654a2.307 2.307 0 0 0-3.262 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.964 42.022V25.434a4.614 4.614 0 0 0-1.351-3.263L18.233 10.79m1.196 31.232V25.434a4.614 4.614 0 0 0-1.352-3.263l-5.613-5.613"/>`),
		g.Group(children),
	)
}