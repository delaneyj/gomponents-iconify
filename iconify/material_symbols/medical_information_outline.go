package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalInformationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18h2v-2h2v-2H9v-2H7v2H5v2h2v2Zm6-3.5h6V13h-6v1.5Zm0 3h4V16h-4v1.5ZM4 22q-.825 0-1.413-.588T2 20V9q0-.825.588-1.413T4 7h5V4q0-.825.588-1.413T11 2h2q.825 0 1.413.588T15 4v3h5q.825 0 1.413.588T22 9v11q0 .825-.588 1.413T20 22H4Zm0-2h16V9h-5q0 .825-.588 1.413T13 11h-2q-.825 0-1.413-.588T9 9H4v11Zm7-11h2V4h-2v5Zm1 5.5Z"/>`),
		g.Group(children),
	)
}