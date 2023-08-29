package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Locker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 14.24a9 9 0 0 1 9 9h0a9 9 0 1 1-9-9Zm0 4a3.8 3.8 0 0 1 .91 7.5v2.4a.91.91 0 1 1-1.82 0v-2.4a3.8 3.8 0 0 1 .91-7.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.2 4.86L6.69 11.25V27C6.69 35.44 24 43.5 24 43.5S41.31 35.44 41.31 27V11.25L25.8 4.86a4.68 4.68 0 0 0-3.6 0Z"/>`),
		g.Group(children),
	)
}