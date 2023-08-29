package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkWifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21L0 9q2.4-2.45 5.5-3.725T12 4q3.425 0 6.525 1.275T24 9L12 21ZM4.35 10.5q1.625-1.175 3.563-1.838T12 8q2.15 0 4.088.663T19.65 10.5l1.45-1.45q-1.95-1.475-4.263-2.263T12 6q-2.525 0-4.838.788T2.9 9.05l1.45 1.45Z"/>`),
		g.Group(children),
	)
}