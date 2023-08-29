package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicineBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMedicineBottle0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M34 10H14a2 2 0 0 0-2 2v30a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V12a2 2 0 0 0-2-2Z"/><path stroke-linecap="round" d="M12 18h24"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 15v6m24-6v6"/><path fill="#555" stroke-linejoin="round" d="M32 4H16v6h16V4Z"/><path stroke-linecap="round" d="M20 31h8m-4-4v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMedicineBottle0)"/>`),
		g.Group(children),
	)
}