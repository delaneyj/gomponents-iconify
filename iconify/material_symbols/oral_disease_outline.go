package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OralDiseaseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22v-9h2V8.4L3.6 5l4-4L9 2.4L6.4 5L9 7.6V13h2v9H5Zm8 0v-9h2V9.875q-1.3-.35-2.15-1.4T12 6q0-1.65 1.175-2.825T16 2q1.65 0 2.825 1.175T20 6q0 1.425-.85 2.475T17 9.875V13h2v9h-6Zm3-14q.825 0 1.413-.588T18 6q0-.825-.588-1.413T16 4q-.825 0-1.413.588T14 6q0 .825.588 1.413T16 8ZM7 20h2v-5H7v5Zm8 0h2v-5h-2v5Zm-8 0h2h-2Zm8 0h2h-2Z"/>`),
		g.Group(children),
	)
}