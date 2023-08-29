package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SolarPower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4V2h3v2H3ZM2 22h9v-4H2.8L2 22ZM6.125 9.325L4.7 7.925L6.825 5.8L8.25 7.2L6.125 9.325ZM3.2 16H11v-4H4l-.8 4ZM12 7Q9.925 7 8.462 5.537T7 2h10q0 2.075-1.463 3.538T12 7Zm-1 4V8h2v3h-2Zm2 11h9l-.8-4H13v4Zm0-6h7.8l-.8-4h-7v4Zm4.875-6.675l-2.1-2.125l1.4-1.4L19.3 7.9l-1.425 1.425ZM18 4V2h3v2h-3Z"/>`),
		g.Group(children),
	)
}