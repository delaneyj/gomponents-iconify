package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryWorkingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBatteryWorkingOne0"><g fill="none"><rect width="36" height="20" x="4" y="14" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><path fill="#fff" d="M42 20h2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2v-8Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 21v6m6-6v6m6-6v6m6-3v3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBatteryWorkingOne0)"/>`),
		g.Group(children),
	)
}