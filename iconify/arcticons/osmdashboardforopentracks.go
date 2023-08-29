package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Osmdashboardforopentracks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.69 11.34L13.28 26L24 27l3.9 9.74Zm-30.2 4.5L32.01 5.5m.83 15.1l9.67 5.41m-25.92-14.5l10.15 5.67m-.14 16.26l-7.88 9.06m17.44-20.04l-6.02 6.92"/>`),
		g.Group(children),
	)
}