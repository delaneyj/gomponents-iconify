package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 12V38H20H28H30M18 10H20H36V24V27"/><path d="M44 14L36 20.75V27.25L44 34V14Z" clip-rule="evenodd"/><path d="M44 44L4 4"/></g>`),
		g.Group(children),
	)
}