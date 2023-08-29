package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindDirectionSe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path d="M0 767q0-156 60.5-298.5t163-244.5t244-163T765 0q157 0 299 61t245 163.5T1473 469t61 297q0 156-61 298.5T1309 1310t-245 163.5t-298 60.5q-155 0-297-60.5t-244.5-163T61 1066T0 767zm168 0q0 162 80.5 299.5t218.5 217t299 79.5q121 0 232-47.5t191.5-127.5t128-190.5T1365 766q0-244-178-422q-178-176-422-176q-121 0-231.5 48T343 344T215.5 535T168 767zm207-8q0-10 10-14l231-104q10-5 15-15l103-229q9-12 16-12q8 0 11 12l248 605q2 11-2 15t-16 2L385 771q-10-3-10-12z" fill="currentColor"/>`),
		g.Group(children),
	)
}