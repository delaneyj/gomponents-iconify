package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 0 0 8.8 39.2l6.91-6.9A16.61 16.61 0 1 1 39.2 8.8A21.43 21.43 0 0 0 24 2.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 45.5A21.5 21.5 0 0 0 39.2 8.8l-6.91 6.9A16.61 16.61 0 1 1 8.8 39.2A21.43 21.43 0 0 0 24 45.5Z"/>`),
		g.Group(children),
	)
}