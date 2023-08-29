package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddAPhotoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 13Zm-8 8q-.825 0-1.413-.588T1 19V7q0-.825.588-1.413T3 5h3.15L8 3h6v2H8.875L7.05 7H3v12h16v-9h2v9q0 .825-.588 1.413T19 21H3ZM19 7V5h-2V3h2V1h2v2h2v2h-2v2h-2Zm-8 10.5q1.875 0 3.188-1.313T15.5 13q0-1.875-1.313-3.188T11 8.5q-1.875 0-3.188 1.313T6.5 13q0 1.875 1.313 3.188T11 17.5Zm0-2q-1.05 0-1.775-.725T8.5 13q0-1.05.725-1.775T11 10.5q1.05 0 1.775.725T13.5 13q0 1.05-.725 1.775T11 15.5Z"/>`),
		g.Group(children),
	)
}