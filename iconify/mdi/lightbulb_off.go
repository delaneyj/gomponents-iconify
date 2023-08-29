package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C9.76 2 7.78 3.05 6.5 4.68l9.81 9.82C17.94 13.21 19 11.24 19 9a7 7 0 0 0-7-7M3.28 4L2 5.27L5.04 8.3C5 8.53 5 8.76 5 9c0 2.38 1.19 4.47 3 5.74V17a1 1 0 0 0 1 1h5.73l4 4L20 20.72L3.28 4M9 20v1a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-1H9Z"/>`),
		g.Group(children),
	)
}