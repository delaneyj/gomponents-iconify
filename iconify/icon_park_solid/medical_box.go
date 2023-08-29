package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMedicalBox0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M41 17H7a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h34a2 2 0 0 0 2-2V19a2 2 0 0 0-2-2ZM34 7H14v10h20V7Z"/><path stroke="#000" stroke-linecap="round" d="M19 29h10m-5-5v10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMedicalBox0)"/>`),
		g.Group(children),
	)
}