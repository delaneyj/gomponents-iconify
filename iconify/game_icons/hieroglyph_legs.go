package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HieroglyphLegs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M166.4 48.07C174.7 139.7 83.71 316.3 30.34 410.6c-13.4 23.7-10.14 47.1 8.03 53.3l173.73-.5c1.7-9.1 1.7-19.1-7.3-31.3c-32.1-16-76.5-6.2-95.6-30.8c-13.33-17.2 84.7-149.3 120.4-259.7c11.3 79.1 32.1 172.1 76.3 262.7c-1.1 23.6-8.8 53.3 9.4 59.6l173.7-.5c1.7-9.1 1.7-19.1-7.3-31.3c-32.1-16.1-85.6-5.7-103.8-31.1c-70.4-98.3-71.2-243.9-99.5-352.93z"/>`),
		g.Group(children),
	)
}