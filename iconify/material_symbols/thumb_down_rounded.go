package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbDownRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16q-.8 0-1.4-.6T1 14v-2q0-.175.037-.375t.113-.375l3-7.05q.225-.5.75-.85T6 3h8q.825 0 1.413.587T16 5v10.175q0 .4-.163.763t-.437.637l-5.425 5.4q-.375.35-.887.425t-.988-.175q-.475-.25-.687-.7t-.088-.925L8.45 16H3ZM20 3q.825 0 1.413.588T22 5v9q0 .825-.588 1.413T20 16q-.825 0-1.413-.588T18 14V5q0-.825.588-1.413T20 3Z"/>`),
		g.Group(children),
	)
}