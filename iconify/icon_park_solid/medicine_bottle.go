package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicineBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMedicineBottle0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M34 10H14a2 2 0 0 0-2 2v30a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V12a2 2 0 0 0-2-2Z"/><path stroke="#000" stroke-linecap="round" d="M12 18h24"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M12 15v6m24-6v6"/><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M32 4H16v6h16V4Z"/><path stroke="#000" stroke-linecap="round" d="M20 31h8m-4-4v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMedicineBottle0)"/>`),
		g.Group(children),
	)
}