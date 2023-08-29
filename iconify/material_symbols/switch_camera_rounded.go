package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchCameraRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h3.15L8.4 3.65q.275-.3.663-.475T9.874 3h4.25q.425 0 .813.175t.662.475L16.85 5H20q.825 0 1.413.588T22 7v12q0 .825-.588 1.413T20 21H4Zm4.85-6.95h6.3l-.85.85q-.275.275-.275.7t.275.7q.275.275.7.275t.7-.275l2.6-2.6q.3-.3.3-.7t-.3-.7l-2.6-2.6q-.275-.275-.687-.287T14.3 9.7q-.275.275-.288.688t.263.712l.925.95H8.8l.925-.95q.275-.275.275-.687T9.7 9.7q-.275-.275-.7-.275t-.7.275l-2.6 2.6q-.3.3-.3.7t.3.7l2.6 2.6q.275.275.7.275t.7-.275q.275-.275.275-.7t-.275-.7l-.85-.85Z"/>`),
		g.Group(children),
	)
}