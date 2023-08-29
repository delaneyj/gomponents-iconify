package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMedicalMark0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M16 15a3 3 0 0 1 3-3h18a3 3 0 0 1 3 3v18a3 3 0 0 1-3 3H19a3 3 0 0 1-3-3V15Z"/><path stroke="#fff" stroke-linecap="round" d="M8 4v40m0-25h8M8 29h8"/><path stroke="#000" stroke-linecap="round" d="M22 24h12m-6-6v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMedicalMark0)"/>`),
		g.Group(children),
	)
}