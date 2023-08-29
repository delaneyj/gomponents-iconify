package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedtimeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.15 18.125L5.875 3.85q1.075-1.1 2.438-1.825T11.25 1q-.45 2.475.275 4.838t2.5 4.137q1.775 1.775 4.138 2.5T23 12.75q-.275 1.575-1.012 2.938t-1.838 2.437Zm-.775 4.925l-2.7-2.7q-.85.325-1.738.488T13.1 21q-2.1 0-3.937-.8t-3.2-2.163Q4.6 16.675 3.8 14.837T3 10.9q0-.95.163-1.838t.487-1.737L.975 4.65L2.4 3.225l18.4 18.4l-1.425 1.425Z"/>`),
		g.Group(children),
	)
}