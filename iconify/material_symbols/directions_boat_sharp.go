package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsBoatSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.95 19l-2.175-7.625L4 10.6V4h5V1h6v3h5v6.6l2.225.775L20.05 19q-1.25 0-2.275-.588T16 17q-.75.825-1.775 1.413T12 19q-1.2 0-2.225-.588T8 17q-.75.825-1.775 1.413T3.95 19ZM2 23v-2h2q1.05 0 2.075-.325T8 19.7q.9.65 1.925.95t2.075.3q1.05 0 2.075-.3T16 19.7q.9.65 1.925.975T20 21h2v2h-2q-1.05 0-2.05-.25T16 22q-.95.5-1.963.75T12 23q-1.025 0-2.038-.25T8 22q-.95.5-1.95.75T4 23H2ZM6 9.95L12 8l6 1.95V6H6v3.95Z"/>`),
		g.Group(children),
	)
}