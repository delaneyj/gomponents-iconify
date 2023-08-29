package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalFiles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMedicalFiles0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M23 42H9a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2v11.5"/><path fill="#555" stroke-linejoin="round" d="M36.636 27C39.046 27 41 28.88 41 31.2c0 3.02-2.91 5.6-4.364 7c-.97.933-2.181 1.867-3.636 2.8c-1.454-.933-2.667-1.867-3.636-2.8c-1.455-1.4-4.364-3.98-4.364-7c0-2.32 1.954-4.2 4.364-4.2c1.517 0 2.854.746 3.636 1.878A4.406 4.406 0 0 1 36.636 27Z"/><path stroke-linecap="round" d="M15 14h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMedicalFiles0)"/>`),
		g.Group(children),
	)
}