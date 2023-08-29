package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImmoscoutTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5L5.879 17.547h2.024V43.5h32.194V17.547h2.024Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.19 27.082a5.18 5.18 0 0 1 5.045-5.24a5.215 5.215 0 0 1 3.687 8.928c-2.134 1.747-8.733 6.793-8.733 6.793h10.286m11.395 0V22.037l-8.345 10.48h10.286M7.903 17.547h32.194"/>`),
		g.Group(children),
	)
}