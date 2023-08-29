package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AllTools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.294 43.5c6.157-22.318 11.74-30.939 19.74-38.619c.912-.875 2.413-.136 2.27 1.12c-1.244 10.878 5.59 29.465 9.402 37.499"/>`),
		g.Group(children),
	)
}