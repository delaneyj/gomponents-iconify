package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5.50934 19.0293C3.08128 15.2358 2.82341 7.06076 10.5522 6.02314C15.8552 5.45972 23.1035 15.1207 24 42C24.8965 15.1207 32.1448 5.45972 37.4478 6.02314C45.1766 7.06076 44.9187 15.2358 42.4907 19.0293"/>`),
		g.Group(children),
	)
}