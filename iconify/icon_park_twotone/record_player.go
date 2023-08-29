package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordPlayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRecordPlayer0"><g fill="none"><rect width="38" height="32" x="5" y="8" stroke="#fff" stroke-width="4" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 8v32"/><circle cx="28" cy="24" r="9" fill="#555" stroke="#fff" stroke-width="4"/><circle cx="28" cy="24" r="3" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 16h8m-8 8h8m-8 8h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRecordPlayer0)"/>`),
		g.Group(children),
	)
}