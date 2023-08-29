package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plexus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.474 8.508l35.051 30.984m0-30.984L6.474 39.492M43.5 11.082l-39 25.836m39 0l-39-25.836"/>`),
		g.Group(children),
	)
}