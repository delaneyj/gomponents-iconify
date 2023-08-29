package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindDirectionE(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path d="M0 767q0-156 60.5-298.5t163-244.5t244-163T765 0q157 0 299 61t245 163.5T1473 469t61 297q0 156-61 298.5T1309 1310t-245 163.5t-298 60.5q-155 0-297-60.5t-244.5-163T61 1066T0 767zm168 0q0 162 80.5 299.5t218.5 217t299 79.5q121 0 232-47.5t191.5-127.5t128-190.5T1365 766q0-244-178-422q-178-176-422-176q-121 0-231.5 48T343 344T215.5 535T168 767zm323 248l90-239q5-10 0-21l-90-235q-5-10 1-16t17-1l598 252q11 2 11 11q0 8-11 10l-598 257q-11 4-17-1.5t-1-16.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}