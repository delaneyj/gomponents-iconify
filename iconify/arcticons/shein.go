package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shein(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.553 39.112c2.44 2.925 5.37 4.388 9.763 4.388h5.857c5.37 0 9.762-4.388 9.762-9.75S32.542 24 27.173 24h-6.345c-5.37 0-9.763-4.388-9.763-9.75s4.393-9.75 9.763-9.75h5.857c4.393 0 7.321.975 9.762 4.388"/>`),
		g.Group(children),
	)
}