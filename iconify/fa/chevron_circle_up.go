package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m1165 1011l102-102q19-19 19-45t-19-45L813 365q-19-19-45-19t-45 19L269 819q-19 19-19 45t19 45l102 102q19 19 45 19t45-19l307-307l307 307q19 19 45 19t45-19zm371-243q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768z"/>`),
		g.Group(children),
	)
}