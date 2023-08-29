package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindDirectionNe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path d="M0 767q0-156 60.5-298.5t163-244.5t244-163T765 0q157 0 299 61t245 163.5T1473 469t61 297q0 156-61 298.5T1309 1310t-245 163.5t-298 60.5q-155 0-297-60.5t-244.5-163T61 1066T0 767zm168 0q0 162 80.5 299.5t218.5 217t299 79.5q121 0 232-47.5t191.5-127.5t128-190.5T1365 766q0-244-178-422q-178-176-422-176q-121 0-231.5 48T343 344T215.5 535T168 767zm267 8q0-7 12-10l604-250q11-2 16.5 3.5t1.5 16.5l-248 604q-2 11-11 11q-11 0-16-11L691 909q-9-14-14-15L447 790q-12-5-12-15z" fill="currentColor"/>`),
		g.Group(children),
	)
}