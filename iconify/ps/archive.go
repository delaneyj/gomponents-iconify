package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 3H43Q25 3 12.5 15.5T0 45v86q0 9 6 15t15 6h22v235q0 17 12.5 29.5T85 429h342q17 0 29.5-12.5T469 387V152h22q9 0 15-6t6-15V45q0-17-12.5-29.5T469 3zM85 387V152h342v235H85zm384-278H43V45h426v64zm-140 90q-32 24-73 24t-73-24q-6-5-15.5-3.5T154 203q-5 7-4 16.5t8 13.5q46 32 98 32t98-32q7-4 8-13.5t-4-16.5q-4-6-13.5-7.5T329 199z"/>`),
		g.Group(children),
	)
}