package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChargingStationLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M133 124.64a6 6 0 0 1 .6 5.59l-16 40a6 6 0 1 1-11.14-4.46L119.14 134H96a6 6 0 0 1-5.57-8.23l16-40a6 6 0 0 1 11.14 4.46L104.86 122H128a6 6 0 0 1 5 2.64Zm113-38V168a22 22 0 0 1-44 0v-40a10 10 0 0 0-10-10h-18v92h18a6 6 0 0 1 0 12H32a6 6 0 0 1 0-12h18V56a22 22 0 0 1 22-22h80a22 22 0 0 1 22 22v50h18a22 22 0 0 1 22 22v40a10 10 0 0 0 20 0V86.63a9.93 9.93 0 0 0-2.93-7.07l-19.31-19.32a6 6 0 0 1 8.48-8.48l19.32 19.31A21.88 21.88 0 0 1 246 86.63ZM162 210V56a10 10 0 0 0-10-10H72a10 10 0 0 0-10 10v154Z"/>`),
		g.Group(children),
	)
}