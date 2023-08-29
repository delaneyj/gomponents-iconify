package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlagiarismOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.45 17q.45 0 .888-.113t.812-.337l1.75 1.75q.275.275.7.275t.7-.275q.275-.275.275-.7t-.275-.7l-1.75-1.75q.225-.375.338-.812T15 13.45Q15 12 13.975 11T11.5 10q-1.45 0-2.475 1.025T8 13.5q0 1.45 1 2.475T11.45 17Zm.05-2q-.625 0-1.062-.438T10 13.5q0-.625.438-1.063T11.5 12q.625 0 1.063.438T13 13.5q0 .625-.438 1.063T11.5 15ZM6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V20q0 .825-.588 1.413T18 22H6Zm7-14V4H6v16h12V9h-4q-.425 0-.713-.288T13 8ZM6 4v5v-5v16V4Z"/>`),
		g.Group(children),
	)
}