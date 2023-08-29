package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m909 1267l102-102q19-19 19-45t-19-45L704 768l307-307q19-19 19-45t-19-45L909 269q-19-19-45-19t-45 19L365 723q-19 19-19 45t19 45l454 454q19 19 45 19t45-19zm627-499q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768z"/>`),
		g.Group(children),
	)
}