package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExtensionOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.8 21H5q-.825 0-1.413-.588T3 19v-3.8q1.2 0 2.1-.762T6 12.5q0-1.175-.9-1.937T3 9.8V6q0-.825.588-1.413T5 4h4q0-1.05.725-1.775T11.5 1.5q1.05 0 1.775.725T14 4h4q.825 0 1.413.588T20 6v4q1.05 0 1.775.725T22.5 12.5q0 1.05-.725 1.775T20 15v4q0 .825-.588 1.413T18 21h-3.8q0-1.25-.787-2.125T11.5 18q-1.125 0-1.913.875T8.8 21ZM5 19h2.125q.6-1.65 1.925-2.325T11.5 16q1.125 0 2.45.675T15.875 19H18v-6h2q.2 0 .35-.15t.15-.35q0-.2-.15-.35T20 12h-2V6h-6V4q0-.2-.15-.35t-.35-.15q-.2 0-.35.15T11 4v2H5v2.2q1.35.5 2.175 1.675T8 12.5q0 1.425-.825 2.6T5 16.8V19Zm7.75-7.75Z"/>`),
		g.Group(children),
	)
}