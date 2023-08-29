package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficConeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M188.52 160h-121l22.22-64h76.52Z" opacity=".2"/><path d="M232 208h-18.31L153.42 34.75A16 16 0 0 0 138.31 24h-20.62a16 16 0 0 0-15.11 10.74L42.31 208H24a8 8 0 0 0 0 16h208a8 8 0 0 0 0-16ZM117.69 40h20.62L155 88h-54Zm-22.26 64h65.14l16.7 48H78.73ZM59.25 208l13.92-40h109.66l13.92 40Z"/></g>`),
		g.Group(children),
	)
}