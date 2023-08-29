package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkWifiThreeBarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.8 11.95q1.325-.95 2.9-1.487t3.3-.538q1.725 0 3.3.537t2.9 1.488l2.9-2.9q-1.95-1.475-4.263-2.263T12 6q-2.525 0-4.838.788T2.9 9.05l2.9 2.9Zm6.2 8.625q-.2 0-.375-.062T11.3 20.3L.7 9.7q-.3-.3-.288-.712T.725 8.3q2.35-2.125 5.238-3.213T12 4q3.175 0 6.063 1.088T23.274 8.3q.3.275.313.688T23.3 9.7L12.7 20.3q-.15.15-.325.213t-.375.062Z"/>`),
		g.Group(children),
	)
}