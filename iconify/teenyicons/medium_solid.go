package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediumSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5v-12ZM4 5H3V4h1.5a.5.5 0 0 1 .4.2l2.6 3.467l2.593-3.458A.5.5 0 0 1 10.5 4H12v1h-1v5h1v1H9v-1h1V6L7.5 9.333L5 6v4h1v1H3v-1h1V5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}