package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Perfume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1006.86 404q18 18 18 59.5t-18 59.5q-8 9-22 25.5t-51.5 71.5t-68.5 116q-32 63-32 96q1 63 64 192h-767q7-14 17.5-37t28.5-74.5t18-80.5q0-16-5.5-34t-10-27.5t-16.5-34.5q-31-60-67-113.5T39.36 546t-20.5-23q-18-18-18-59.5t18-59.5l431-125v-87q-27 0-45.5-18.5t-18.5-45.5q-86 0-130 21q-8 46-43.5 76.5t-82.5 30.5q-53 0-90.5-37.5T1.86 128t37.5-90.5T129.86 0q39 0 71.5 22.5t46.5 57.5q50-15 135-16l1 .5l2 .5v-1q0-26 18.5-45t45.5-19h127q27 0 45.5 19t18.5 45v64q0 27-18.5 45.5t-45.5 18.5v88z"/>`),
		g.Group(children),
	)
}