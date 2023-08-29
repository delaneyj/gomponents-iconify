package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddLocationAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12q.825 0 1.413-.588T14 10q0-.825-.588-1.413T12 8q-.825 0-1.413.588T10 10q0 .825.588 1.413T12 12Zm0-10q.5 0 1 .063t1 .187V4q0 .825.588 1.413T16 6h1v1q0 .825.588 1.413T19 9h.925q.05.275.063.588T20 10.2q0 1.65-.763 3.275t-1.824 3.05q-1.063 1.425-2.2 2.538t-1.888 1.762q-.275.25-.625.375t-.7.125q-.35 0-.7-.125t-.625-.375Q9.05 19.325 7.8 17.9t-2.087-2.762q-.838-1.338-1.275-2.575T4 10.2q0-3.75 2.413-5.975T12 2Zm6 3h-2q-.425 0-.713-.288T15 4q0-.425.288-.713T16 3h2V1q0-.425.288-.713T19 0q.425 0 .713.288T20 1v2h2q.425 0 .713.288T23 4q0 .425-.288.713T22 5h-2v2q0 .425-.288.713T19 8q-.425 0-.713-.288T18 7V5Z"/>`),
		g.Group(children),
	)
}