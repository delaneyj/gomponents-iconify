package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VipCrownTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.806 5.2L7.005 8l4.186-5.86a1 1 0 0 1 1.628-.001L17.005 8l4.2-2.799a1 1 0 0 1 1.547.95L21.11 20.116a1 1 0 0 1-.993.883H3.894a1 1 0 0 1-.993-.883L1.258 6.149A1 1 0 0 1 2.806 5.2Zm9.2 9.8a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}