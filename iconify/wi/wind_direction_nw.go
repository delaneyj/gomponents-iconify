package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindDirectionNw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path d="M0 767q0-156 60.5-298.5t163-244.5t244-163T765 0q157 0 299 61t245 163.5T1473 469t61 297q0 156-61 298.5T1309 1310t-245 163.5t-298 60.5q-155 0-297-60.5t-244.5-163T61 1066T0 767zm168 0q0 162 80.5 299.5t218.5 217t299 79.5q121 0 232-47.5t191.5-127.5t128-190.5T1365 766q0-244-178-422q-178-176-422-176q-121 0-231.5 48T343 344T215.5 535T168 767zm350-229q-2-11 3.5-17t15.5-2l602 249q12 2 12 10q0 9-12 14L910 896q-10 4-15 15l-104 230q-5 11-14 11q-8 0-10-11z" fill="currentColor"/>`),
		g.Group(children),
	)
}