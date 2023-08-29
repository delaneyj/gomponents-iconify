package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NordlockerCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.59 36.89c-7.003-9.732-4.792-23.299 4.94-30.303A21.71 21.71 0 0 1 24 2.5c11.99.117 21.614 9.93 21.497 21.92a21.71 21.71 0 0 1-4.087 12.47L31.08 20l-1.86 3.17l1.88 3.23L24 14.17l-5.27 9l1.9 3.27L16.91 20L6.59 36.89Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.595 38.858h7.993c4.338 0 3.612-6.529-.739-5.081c0-3.626-6.529-3.626-6.529.726c-3.629-.726-3.629 4.355-.726 4.355Z"/>`),
		g.Group(children),
	)
}