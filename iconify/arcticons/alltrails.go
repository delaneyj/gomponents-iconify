package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alltrails(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m41.82 30.535l-8.136-11.784l-3.989 2.336l-5.905-9.79L4.5 35.784c18.136-12.337 27.782-3.774 39 .919"/>`),
		g.Group(children),
	)
}