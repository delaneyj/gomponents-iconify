package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkWifiThreeBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21L0 9q2.4-2.45 5.5-3.725T12 4q3.425 0 6.525 1.275T24 9L12 21Zm-6.2-9.05q1.325-.95 2.9-1.487t3.3-.538q1.725 0 3.3.537t2.9 1.488l2.9-2.9q-1.95-1.475-4.263-2.263T12 6q-2.525 0-4.838.788T2.9 9.05l2.9 2.9Z"/>`),
		g.Group(children),
	)
}