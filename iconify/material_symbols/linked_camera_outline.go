package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkedCameraOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13Zm8.65-6q0-1.95-1.35-3.3T16 2.35V1q2.5 0 4.25 1.75T22 7h-1.35ZM18 7q0-.825-.588-1.413T16 5V3.65q1.375 0 2.337.975T19.35 7H18ZM4 21q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h3.15L9 3h6v2H9.875L8.05 7H4v12h16V8h2v11q0 .825-.588 1.413T20 21H4Zm8-3.5q1.875 0 3.188-1.313T16.5 13q0-1.875-1.313-3.188T12 8.5q-1.875 0-3.188 1.313T7.5 13q0 1.875 1.313 3.188T12 17.5Zm0-2q-1.05 0-1.775-.725T9.5 13q0-1.05.725-1.775T12 10.5q1.05 0 1.775.725T14.5 13q0 1.05-.725 1.775T12 15.5Z"/>`),
		g.Group(children),
	)
}