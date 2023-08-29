package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SceneOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 22V7q0-.825-.588-1.413T18 5h-1v1.3q0 .3-.2.5t-.5.2h-5.6q-.35 0-.562-.375T10.1 5.9L12 1.8q.175-.375.513-.588T13.3 1h2.3q.6 0 1 .45T17 2.5V3h1q1.65 0 2.825 1.175T22 7v15h-2ZM12.7 5H15V3h-1.4l-.9 2ZM5 22q-1.275 0-2.138-.863T2 19v-2.5q0-.825.55-1.538T4 14.1V12q0-.825.588-1.413T6 10h8q.825 0 1.413.588T16 12v2.1q.9.15 1.45.825T18 16.5V19q0 1.275-.863 2.138T15 22H5Zm1-10v2.5q.45.375.725.888T7 16.5v.5h6v-.5q0-.6.275-1.113T14 14.5V12H6Zm-1 8h10q.45 0 .725-.313T16 19v-2.5q0-.225-.138-.363T15.5 16q-.225 0-.363.138T15 16.5V19H5v-2.5q0-.225-.138-.363T4.5 16q-.225 0-.363.138T4 16.5V19q0 .375.275.688T5 20Zm8-3H7h6Zm-7-5h8h-8Zm-1 8h10H5Z"/>`),
		g.Group(children),
	)
}