package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm486 1493h147q28 0 44-19.5t31-55.5l447-1044q11-27-15-51t-59-24h-149q-28 0-44 19.5t-31 55.5l-298 695l-298-695q-10-24-17.5-37.5T557 311t-34-12H374q-33 0-59.5 24T299 374l448 1044q15 36 31 55.5t44 19.5z"/>`),
		g.Group(children),
	)
}