package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashingMachineOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWashingMachineOne0"><g fill="none"><rect width="32" height="40" x="8" y="4" stroke="#fff" stroke-width="4" rx="2"/><path fill="#555" stroke="#fff" stroke-width="4" d="M8 12a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H10a2 2 0 0 0-2 2v6Z"/><circle cx="14" cy="9" r="2" fill="#fff"/><circle cx="20" cy="9" r="2" fill="#fff"/><circle cx="24" cy="29" r="7" fill="#555" stroke="#fff" stroke-width="4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWashingMachineOne0)"/>`),
		g.Group(children),
	)
}