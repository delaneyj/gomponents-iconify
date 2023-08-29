package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 2C10.477 2 6 6.477 6 12a9.978 9.978 0 0 0 3.365 7.482c.343.304.561.645.64.986L10.82 24h10.36l.815-3.532c.079-.34.297-.682.64-.986A9.978 9.978 0 0 0 26 12c0-5.523-4.477-10-10-10Zm4.719 24H11.28l.297 1.287A3.5 3.5 0 0 0 14.988 30h2.023a3.5 3.5 0 0 0 3.41-2.713L20.72 26Z"/>`),
		g.Group(children),
	)
}