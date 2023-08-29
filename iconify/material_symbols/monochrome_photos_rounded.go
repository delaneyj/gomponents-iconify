package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonochromePhotosRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17.5q1.875 0 3.188-1.313T16.5 13q0-1.875-1.313-3.188T12 8.5q-1.875 0-3.188 1.313T7.5 13q0 1.875 1.313 3.188T12 17.5Zm0-2q-1.05 0-1.775-.725T9.5 13q0-1.05.725-1.775T12 10.5q1.05 0 1.775.725T14.5 13q0 1.05-.725 1.775T12 15.5ZM4 21q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h3.15L8.4 3.65q.275-.3.663-.475T9.874 3h4.25q.425 0 .813.175t.662.475L16.85 5H20q.825 0 1.413.588T22 7v12q0 .825-.588 1.413T20 21H4Zm8-2h8V7h-4.05l-1.825-2H12v14Zm-8 0V7h4.05l1.825-2h4.25l1.825 2H20v12H4Z"/>`),
		g.Group(children),
	)
}