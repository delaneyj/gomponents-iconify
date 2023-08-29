package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FollowTheSigns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 5.5q-.825 0-1.413-.588T7.5 3.5q0-.825.588-1.413T9.5 1.5q.825 0 1.413.588T11.5 3.5q0 .825-.588 1.413T9.5 5.5ZM3 23L5.75 8.9L4 9.65V13H2V8.3l5.25-2.15q.675-.275 1.375-.063T9.7 6.95l.95 1.6q.675 1.1 1.812 1.775T15 11v2q-1.65 0-3.063-.7T9.55 10.4l-.6 3L11 15.45V23H9v-6l-2.15-2l-1.75 8H3Zm13.75 0V9H13V2h9v7h-3.75v14h-1.5Zm1.275-15.025L20.5 5.5l-2.475-2.475L16.95 4.1l.65.65h-3.1v1.5h3.1l-.65.65l1.075 1.075Z"/>`),
		g.Group(children),
	)
}