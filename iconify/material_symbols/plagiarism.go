package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plagiarism(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.45 17q.45 0 .888-.113t.812-.337L15.6 19l1.4-1.4l-2.45-2.45q.225-.375.338-.812T15 13.45Q15 12 13.975 11T11.5 10q-1.45 0-2.475 1.025T8 13.5q0 1.45 1 2.475T11.45 17Zm.05-2q-.625 0-1.062-.438T10 13.5q0-.625.438-1.063T11.5 12q.625 0 1.063.438T13 13.5q0 .625-.438 1.063T11.5 15ZM6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h8l6 6v12q0 .825-.588 1.413T18 22H6Zm7-13h5l-5-5v5Z"/>`),
		g.Group(children),
	)
}