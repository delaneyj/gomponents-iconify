package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicineChest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMedicineChest0"><g fill="none"><rect width="38" height="26" x="5" y="16" fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="3"/><path fill="#fff" d="M19 8h10V4H19v4Zm11 1v7h4V9h-4Zm-12 7V9h-4v7h4Zm11-8a1 1 0 0 1 1 1h4a5 5 0 0 0-5-5v4ZM19 4a5 5 0 0 0-5 5h4a1 1 0 0 1 1-1V4Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 29h12m-6-6v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMedicineChest0)"/>`),
		g.Group(children),
	)
}