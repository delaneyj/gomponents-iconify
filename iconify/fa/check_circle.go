package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1284 606q0-28-18-46l-91-90q-19-19-45-19t-45 19L677 877L451 651q-19-19-45-19t-45 19l-91 90q-18 18-18 46q0 27 18 45l362 362q19 19 45 19q27 0 46-19l543-543q18-18 18-45zm252 162q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768z"/>`),
		g.Group(children),
	)
}