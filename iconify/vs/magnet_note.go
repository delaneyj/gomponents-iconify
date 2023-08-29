package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1664 128v640q0 116-.5 170.5T1662 1082t-4.5 132t-9.5 108.5t-15.5 102.5t-22.5 84t-31.5 83t-42.5 72H0q79-242 103.5-447.5T128 640V128h442q-24 60-24 128q0 145 102.5 247.5T896 606t247.5-102.5T1246 256q0-68-24-128h442zM896 0q106 0 181 75t75 181t-75 181t-181 75t-181-75t-75-181t75-181T896 0zm0 128q8 0 14-7t6-17t-6-17t-14-7q-73 0-124.5 51.5T720 256q0 7 8 13.5t16 6.5q10 0 17-6t7-14q0-52 38-90t90-38z"/>`),
		g.Group(children),
	)
}