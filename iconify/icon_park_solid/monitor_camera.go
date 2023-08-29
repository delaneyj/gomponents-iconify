package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMonitorCamera0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 34c8.284 0 15-6.716 15-15c0-8.284-6.716-15-15-15c-8.284 0-15 6.716-15 15c0 8.284 6.716 15 15 15Z"/><path fill="#000" stroke="#000" d="M24 25a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path stroke="#fff" stroke-linecap="round" d="M19.369 34L16 44h16l-3.396-10h-9.235Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" d="M12 44h24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMonitorCamera0)"/>`),
		g.Group(children),
	)
}