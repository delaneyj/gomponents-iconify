package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NyYankees(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.986 39.875s8.921-15.112-.29-29.166l28.258 31.06s-8.237-21.51 2.35-29.842"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.315 42.5L24 25.827S28.456 12.648 32.617 5.5M24 25.827S19.544 12.648 15.383 5.5"/>`),
		g.Group(children),
	)
}