package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaWideAngle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-2.275 0-4.538-.213T3 19q-.525-1.725-.763-3.463T2 12q0-1.8.238-3.538T3 5q1.975-.5 4.2-.75T12 4q2.575 0 4.8.25T21 5q.525 1.725.763 3.463T22 12q0 1.8-.25 3.538T21 19q-2.2.575-4.463.788T12 20Z"/>`),
		g.Group(children),
	)
}