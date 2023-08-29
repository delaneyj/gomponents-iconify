package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefreshAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 7.5A7 7 0 0 1 13 3.17m1.5 4.33A7 7 0 0 1 2 11.83m3-.33H1.5V15m12-15v3.5H10"/>`),
		g.Group(children),
	)
}