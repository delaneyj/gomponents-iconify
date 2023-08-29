package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M2020 11q28 20 28 53v1408q0 20-11 36t-29 23l-640 256q-24 11-48 0l-616-246l-616 246q-10 5-24 5q-19 0-36-11q-28-20-28-53V320q0-20 11-36t29-23L680 5q24-11 48 0l616 246L1960 5q32-13 60 6zM736 146v1270l576 230V376zM128 363v1270l544-217V146zm1792 1066V159l-544 217v1270z"/>`),
		g.Group(children),
	)
}