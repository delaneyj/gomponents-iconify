package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VehicleCableCarTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M25.151 3.007a.75.75 0 1 1 .198 1.487L17.5 5.54v3.71c0 .644-.187 1.245-.51 1.75h2.26A3.75 3.75 0 0 1 23 14.75v3.75H5v-3.75A3.75 3.75 0 0 1 8.75 11h5.5A1.75 1.75 0 0 0 16 9.25V5.74L2.85 7.494a.75.75 0 0 1-.199-1.487l13.35-1.78V3.75a.75.75 0 0 1 1.5 0v.277l7.65-1.02ZM23.001 20v2.25A2.75 2.75 0 0 1 20.25 25H7.75A2.75 2.75 0 0 1 5 22.25V20h18Z"/>`),
		g.Group(children),
	)
}