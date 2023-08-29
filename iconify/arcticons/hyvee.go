package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hyvee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.162 4.5v31.2c0 5.426 4.822 7.22 9.75 7.8V16.875h5.85V43.5c5.107-.284 10.076-1.953 10.076-7.8V4.5H26.762v8.775h-5.85V4.5Z"/>`),
		g.Group(children),
	)
}