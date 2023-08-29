package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestRemoteComfortSensor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-1.65 0-2.825-1.175T8 17q0-1.65 1.175-2.825T12 13q1.65 0 2.825 1.175T16 17q0 1.65-1.175 2.825T12 21ZM3.05 8.6L1.575 7.225q2.025-1.975 4.7-3.1T12 3q3.05 0 5.725 1.125t4.7 3.1L20.95 8.6q-1.775-1.675-4.05-2.638T12 5q-2.625 0-4.9.963T3.05 8.6Zm14.25 3.425q-1.05-.95-2.4-1.488T12 10q-1.55 0-2.888.537T6.7 12.025L5.225 10.65Q6.575 9.4 8.313 8.7T12 8q1.95 0 3.675.7t3.075 1.95l-1.45 1.375Z"/>`),
		g.Group(children),
	)
}