package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateNinetyDegreesCwOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22q-1.875 0-3.513-.713t-2.85-1.924q-1.212-1.213-1.925-2.85T2 13q0-3.75 2.625-6.375T11 4h.15L9.6 2.45L11 1l4 4l-4 4l-1.4-1.45L11.15 6H11Q8.075 6 6.037 8.038T4 13q0 2.925 2.038 4.963T11 20q.875 0 1.725-.213t1.625-.637l1.45 1.45q-1.075.7-2.3 1.05T11 22Zm6-3l-6-6l6-6l6 6l-6 6Zm0-2.85L20.15 13L17 9.85L13.85 13L17 16.15ZM17 13Z"/>`),
		g.Group(children),
	)
}