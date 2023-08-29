package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M0 640q0-130 51-248.5t136.5-204T391.5 51T640 0h768q130 0 248.5 51t204 136.5t136.5 204t51 248.5t-51 248.5t-136.5 204t-204 136.5t-248.5 51H640q-130 0-248.5-51t-204-136.5T51 888.5T0 640zm1408 512q104 0 198.5-40.5T1770 1002t109.5-163.5T1920 640t-40.5-198.5T1770 278t-163.5-109.5T1408 128t-198.5 40.5T1046 278T936.5 441.5T896 640t40.5 198.5T1046 1002t163.5 109.5T1408 1152z"/>`),
		g.Group(children),
	)
}