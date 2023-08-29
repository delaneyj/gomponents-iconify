package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EarthSphereLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12a9.966 9.966 0 0 1-.832 4M12 22a9.966 9.966 0 0 0 7.071-2.929M2 12a9.966 9.966 0 0 1 2.929-7.071M12 2a9.966 9.966 0 0 0-4 .832m0 18.336A10.02 10.02 0 0 1 2.832 16m13-13A10.02 10.02 0 0 1 21 8.168M9 6.803A5.998 5.998 0 0 0 6 12c0 2.22 1.207 4.16 3 5.197M9 6.803A6.003 6.003 0 0 1 17.659 10M9 6.803l1.603 1.708a1.83 1.83 0 0 1-.236 2.714l-.59.442a1.373 1.373 0 0 0 .21 2.327v0c.92.46 1.027 1.729.198 2.336L9 17.197m0 0c.883.51 1.907.803 3 .803c1.792 0 3.4-.786 4.5-2.031M17.659 10A5.99 5.99 0 0 1 18 12a5.977 5.977 0 0 1-1.5 3.969M17.659 10l-1.101.349a2 2 0 0 0-1.252 2.653L16.5 15.97"/>`),
		g.Group(children),
	)
}