package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DroneZone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m10.5 3l-1.03.348A12.41 12.41 0 0 1 5.5 4m-5-1l1.03.348C2.81 3.78 4.15 4 5.5 4m0 0v3.5m18-4.5l-1.03.348A12.41 12.41 0 0 1 18.5 4m-5-1l1.03.348C15.81 3.78 17.15 4 18.5 4m0 0v3.5M5 15.5C3 15.5.5 18 .5 21M19 15.5c2 0 4.5 2.5 4.5 5.5M.5 7.5v4c5.5 0 8.5 6 8.5 6h6s3-6 8.5-6v-4H.5Zm11.5 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}