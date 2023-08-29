package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RingVolumeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.025 21L.4 17.475l1-1.025q2.175-2.225 4.963-3.337T12 12q2.85 0 5.625 1.113T22.6 16.45l1 1.025L19.975 21L16 18v-3.35q-.95-.3-1.95-.475T12 14q-1.05 0-2.05.175T8 14.65V18l-3.975 3ZM11 7V2h2v5h-2Zm6.6 2.85L16.2 8.4l3.55-3.55l1.4 1.45l-3.55 3.55Zm-11.2 0L2.85 6.3l1.4-1.45L7.8 8.4L6.4 9.85Z"/>`),
		g.Group(children),
	)
}