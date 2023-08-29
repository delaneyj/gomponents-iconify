package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBatteryEmpty0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 14H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h32a2 2 0 0 0 2-2V16a2 2 0 0 0-2-2Z"/><path fill="#fff" d="M42 20h2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2v-8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBatteryEmpty0)"/>`),
		g.Group(children),
	)
}