package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Direction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m50 0l-1.455 2.518l-14.063 24.328h7.114v22.63a75.427 75.427 0 0 1 8.4 11.337a77.128 77.128 0 0 1 8.408-11.329V26.846h7.114L50 0zM30.89 34.348L0 37.057l1.666 2.382l16.098 23.034l2.51-5.377c7.137 5.62 17.034 16.617 19.283 35.492c.441-4.589 1.239-9.144 2.505-13.58c1.368-4.927 3.288-9.679 5.665-14.192C41.472 53.731 33.27 46.26 27.373 41.885l3.518-7.537zm38.22 0l3.517 7.537c-10.755 7.975-29.184 26.22-29.459 57.969l16.805.146c.203-23.476 11.728-36.585 19.754-42.904l2.51 5.377L100 37.057l-2.896-.254l-27.995-2.455z" color="currentColor"/>`),
		g.Group(children),
	)
}