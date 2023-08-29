package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M704 1056V480q0-14-9-23t-23-9H416q-14 0-23 9t-9 23v576q0 14 9 23t23 9h256q14 0 23-9t9-23zm448 0V480q0-14-9-23t-23-9H864q-14 0-23 9t-9 23v576q0 14 9 23t23 9h256q14 0 23-9t9-23zm384-288q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768z"/>`),
		g.Group(children),
	)
}