package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleLoadLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.947 8.379c-.075-4.122-3.351-7.412-7.303-7.337c-3.917.073-6.853 2.89-6.83 6.964H.204c-.217.22-.217.876 0 1.096l1.911 1.758a.546.546 0 0 0 .785 0l1.92-1.758c.219-.22.219-.876 0-1.096H3.204C3.18 4.668 5.462 2.461 8.67 2.402c3.232-.062 5.912 2.63 5.975 6.002c.062 3.322-2.445 6.077-5.613 6.216v1.361c3.898-.139 6.991-3.521 6.915-7.602z"/>`),
		g.Group(children),
	)
}