package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func License(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 0 45.5 24A21.51 21.51 0 0 0 24 2.5Zm0 9.77a11.73 11.73 0 0 0-11.45 9.29h5.59a6.35 6.35 0 1 1 0 4.88h-5.6A11.72 11.72 0 1 0 24 12.27Z"/>`),
		g.Group(children),
	)
}