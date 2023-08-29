package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMonitorOne0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 8a2 2 0 0 1 2-2h34a2 2 0 0 1 2 2v24a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V8Z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 26a2 2 0 0 1 2-2h34a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-6Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m22 12l-4 5m10-3l-3 4"/><circle cx="24" cy="29" r="2" fill="#000"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m17 34l-3 8h20l-3-8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMonitorOne0)"/>`),
		g.Group(children),
	)
}