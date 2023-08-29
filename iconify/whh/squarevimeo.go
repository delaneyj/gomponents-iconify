package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarevimeo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-240-768q-46 0-78.5 43.5t-33.5 98.5q16-14 36-14q21 0 32.5 10.5t11.5 37.5q0 25-17.5 63.5t-40.5 67.5t-38 29q-7 0-14-16t-14-47t-12-59t-12.5-70.5t-11.5-63.5q-15-80-64-80q-72 0-144 128l16 32q11-13 26-22.5t22-9.5t12.5 4.5t8.5 13t5 15.5t3.5 17t2.5 14q4 18 12 66t16.5 86t22.5 79.5t37.5 65t55.5 23.5q39 0 104.5-68.5t116.5-161.5t51-154q0-128-112-128z"/>`),
		g.Group(children),
	)
}