package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureKelvin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 5h3v6l5-6h4l-5.12 5.78L19 20h-3.62l-3.62-6.83L10 15.15V20H7V5Z"/>`),
		g.Group(children),
	)
}