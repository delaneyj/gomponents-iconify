package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectronicDoorLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTElectronicDoorLock0"><g fill="none"><rect width="26" height="40" x="6" y="4" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><rect width="24" height="8" x="20" y="30" fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="4"/><circle cx="12" cy="12" r="2" fill="#fff"/><circle cx="12" cy="18" r="2" fill="#fff"/><circle cx="12" cy="24" r="2" fill="#fff"/><circle cx="19" cy="12" r="2" fill="#fff"/><circle cx="19" cy="18" r="2" fill="#fff"/><circle cx="19" cy="24" r="2" fill="#fff"/><circle cx="26" cy="12" r="2" fill="#fff"/><circle cx="26" cy="18" r="2" fill="#fff"/><circle cx="26" cy="24" r="2" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTElectronicDoorLock0)"/>`),
		g.Group(children),
	)
}