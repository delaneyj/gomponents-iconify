package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SolarPowerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2 22l2-10h16l2 10H2ZM3 4V2h3v2H3Zm1.45 16H11v-2H4.85l-.4 2ZM6.125 9.325L4.7 7.925L6.825 5.8L8.25 7.2L6.125 9.325ZM5.25 16H11v-2H5.65l-.4 2ZM12 7Q9.925 7 8.462 5.537T7 2h2q0 1.25.875 2.125T12 5q1.25 0 2.125-.875T15 2h2q0 2.075-1.463 3.538T12 7Zm0-5Zm-1 9V8h2v3h-2Zm2 9h6.55l-.4-2H13v2Zm0-4h5.75l-.4-2H13v2Zm4.875-6.675l-2.1-2.125l1.4-1.4L19.3 7.9l-1.425 1.425ZM18 4V2h3v2h-3Z"/>`),
		g.Group(children),
	)
}