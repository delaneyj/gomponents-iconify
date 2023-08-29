package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tunerly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 14.25V24A9.75 9.75 0 0 0 24 24a9.75 9.75 0 0 1 19.5 0v9.75"/>`),
		g.Group(children),
	)
}