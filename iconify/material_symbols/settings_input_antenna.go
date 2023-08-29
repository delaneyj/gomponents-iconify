package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsInputAntenna(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 12q0-2.3.863-4.3t2.362-3.487q1.5-1.488 3.5-2.35T12 1q2.275 0 4.275.863t3.5 2.35q1.5 1.487 2.363 3.487T23 12h-2q0-1.875-.713-3.513T18.35 5.626Q17.125 4.4 15.487 3.7T12 3q-1.85 0-3.488.7T5.65 5.625Q4.425 6.85 3.712 8.488T3 12H1Zm4 0q0-2.95 2.05-4.975T12 5q2.9 0 4.95 2.025T19 12h-2q0-2.075-1.463-3.538T12 7Q9.925 7 8.462 8.463T7 12H5Zm4 10.4L7.6 21l3.4-3.4v-3.3q-.675-.3-1.088-.925T9.5 12q0-1.05.725-1.775T12 9.5q1.05 0 1.775.725T14.5 12q0 .75-.413 1.375T13 14.3v3.3l3.4 3.4l-1.4 1.4l-3-3l-3 3Z"/>`),
		g.Group(children),
	)
}