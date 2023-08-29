package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitedThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.032 23.2A11.452 11.452 0 0 1 4.5 16C4.5 9.649 9.649 4.5 16 4.5c2.725 0 5.23.948 7.2 2.532L7.032 23.2ZM8.8 24.968L24.968 8.8A11.452 11.452 0 0 1 27.5 16c0 6.351-5.149 11.5-11.5 11.5c-2.725 0-5.23-.948-7.2-2.532ZM2 16c0 7.732 6.268 14 14 14s14-6.268 14-14S23.732 2 16 2S2 8.268 2 16Z"/>`),
		g.Group(children),
	)
}