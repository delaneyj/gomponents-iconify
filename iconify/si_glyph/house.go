package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func House(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m16.794 8.77l-3.81-3.906V2.017l-.968.022v1.728L9.502 1.245a.711.711 0 0 0-1.003 0L1.206 8.771a.713.713 0 0 0 0 1.002a.71.71 0 0 0 1.003-.001L9 2.982l6.793 6.79a.704.704 0 0 0 1.001.001a.715.715 0 0 0 0-1.003z"/><path d="M4.043 9.532v5.69c0 .394.218.786.567.786h2.277l.064-3.993h4.074l-.002 3.993h2.303c.349 0 .632-.391.632-.786V9.531L9 4.625L4.043 9.532z"/></g>`),
		g.Group(children),
	)
}