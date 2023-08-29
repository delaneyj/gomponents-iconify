package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OralDiseaseRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.425 0-.713-.288T5 21v-7q0-.425.288-.713T6 13h1V8.4L5 6.425Q4.425 5.85 4.425 5T5 3.575L6.875 1.7q.3-.3.713-.3t.712.3q.3.3.3.713t-.3.712L6.4 5l2.025 2q.275.275.425.65t.15.775V13h1q.425 0 .713.288T11 14v7q0 .425-.288.713T10 22H6Zm8 0q-.425 0-.713-.288T13 21v-7q0-.425.288-.713T14 13h1V9.875q-1.3-.35-2.15-1.4T12 6q0-1.65 1.175-2.825T16 2q1.65 0 2.825 1.175T20 6q0 1.425-.85 2.475T17 9.875V13h1q.425 0 .713.288T19 14v7q0 .425-.288.713T18 22h-4Zm2-14q.825 0 1.413-.588T18 6q0-.825-.588-1.413T16 4q-.825 0-1.413.588T14 6q0 .825.588 1.413T16 8Z"/>`),
		g.Group(children),
	)
}