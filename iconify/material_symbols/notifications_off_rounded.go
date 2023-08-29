package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationsOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.15 19H5q-.425 0-.713-.288T4 18q0-.425.288-.713T5 17h1v-7q0-.825.213-1.625T6.85 6.85L10 10H7.2L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.288.688t-.288.712q-.275.275-.7.275t-.7-.275L16.15 19ZM18 15.15l-9.8-9.8q.5-.4 1.075-.7T10.5 4.2v-.7q0-.625.438-1.063T12 2q.625 0 1.063.438T13.5 3.5v.7q2 .5 3.25 2.113T18 10v5.15ZM12 22q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Z"/>`),
		g.Group(children),
	)
}