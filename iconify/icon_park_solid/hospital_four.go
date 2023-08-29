package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHospitalFour0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 8a2 2 0 0 1 2-2h20a2 2 0 0 1 2 2v34H6a2 2 0 0 1-2-2V8Z"/><path stroke="#000" d="M21 42v-9a3 3 0 0 0-3-3h-4a3 3 0 0 0-3 3v9"/><path fill="#fff" stroke="#fff" d="M28 24h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H28V24Z"/><path stroke="#000" stroke-linecap="round" d="M12 18h8m14 12h4m-4 6h4M16 14v8"/><path stroke="#fff" stroke-linecap="round" d="M7 42h18"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHospitalFour0)"/>`),
		g.Group(children),
	)
}