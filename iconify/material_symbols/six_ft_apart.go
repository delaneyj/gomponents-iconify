package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixFtApart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 22q-.2 0-.35-.15T7 21.5v-3q0-.2.15-.35T7.5 18H10v1H8v.5h1.5q.2 0 .35.15T10 20v1.5q0 .2-.15.35T9.5 22h-2Zm4 0v-4H14v1h-1.5v.5h1v1h-1V22h-1Zm4 0v-3h-1v-1h3v1h-1v3h-1ZM8 21h1v-.5H8v.5Zm-3-4l-3-3l3-3l1.05 1.05l-.95.95h13.775l-.925-.95L19 11l3 3l-3 3l-1.05-1.05l.95-.95H5.125l.925.95L5 17Zm-3-7v-.575q0-.6.325-1.1t.9-.75q.65-.275 1.338-.425T6 7q.75 0 1.438.15t1.337.425q.575.25.9.75t.325 1.1V10H2Zm12 0v-.575q0-.6.325-1.1t.9-.75q.65-.275 1.338-.425T18 7q.75 0 1.438.15t1.337.425q.575.25.9.75t.325 1.1V10h-8ZM6 6q-.825 0-1.413-.588T4 4q0-.825.588-1.413T6 2q.825 0 1.413.588T8 4q0 .825-.588 1.413T6 6Zm12 0q-.825 0-1.413-.588T16 4q0-.825.588-1.413T18 2q.825 0 1.413.588T20 4q0 .825-.588 1.413T18 6Z"/>`),
		g.Group(children),
	)
}