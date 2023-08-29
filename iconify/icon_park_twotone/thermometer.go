package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thermometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTThermometer0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M18 26V10a6 6 0 0 1 12 0v16a9.985 9.985 0 0 1 4 8c0 5.523-4.477 10-10 10s-10-4.477-10-10a9.985 9.985 0 0 1 4-8Z"/><path stroke-linecap="round" d="M24 17v13"/><path fill="#555" d="M24 38a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTThermometer0)"/>`),
		g.Group(children),
	)
}