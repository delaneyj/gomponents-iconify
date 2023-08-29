package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.49 21.49 0 0 0 7.86 38.19l-4.95 5a1.38 1.38 0 0 0 1 2.36H24a21.5 21.5 0 0 0 0-43Zm0 34.2A12.53 12.53 0 0 1 11.64 24A12.53 12.53 0 0 1 24 11.3A12.53 12.53 0 0 1 36.36 24A12.53 12.53 0 0 1 24 36.7Z"/>`),
		g.Group(children),
	)
}