package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipPreviousBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M22 11v7.967c0 2.31-2.134 3.769-3.87 2.648L7.34 14.646c-1.787-1.154-1.787-4.14 0-5.294l10.79-6.968c1.735-1.12 3.87.34 3.87 2.649V7M2 5v7m0 7v-3"/>`),
		g.Group(children),
	)
}