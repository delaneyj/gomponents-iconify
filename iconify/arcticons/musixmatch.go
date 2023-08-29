package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Musixmatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 31.2l31-24.1c2.4-1.9 6-.2 6 2.9v28c0 3.1-3.5 4.8-6 2.9l-31-24.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.5 31.2l-31-24.1c-2.4-1.9-6-.2-6 2.9v28c0 3.1 3.5 4.8 6 2.9l31-24.1"/>`),
		g.Group(children),
	)
}