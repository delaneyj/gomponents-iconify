package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedroomChildOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17h1.5v-1.5h9V17H18v-4.15q0-.75-.413-1.337T16.5 10.65V9q0-.825-.588-1.413T14.5 7h-5q-.825 0-1.413.588T7.5 9v1.65q-.675.275-1.088.863T6 12.85V17Zm1.5-3v-1.15q0-.35.25-.6t.6-.25h7.3q.35 0 .6.25t.25.6V14h-9ZM9 10.5v-2h6v2H9ZM4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v16q0 .825-.588 1.413T20 22H4Zm0-2h16V4H4v16Zm0 0V4v16Z"/>`),
		g.Group(children),
	)
}