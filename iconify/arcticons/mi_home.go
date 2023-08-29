package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.964 29.122a17.389 17.389 0 0 1-.242-2.9V4.5L24 10.94L41.278 4.5v9.444m0 5.96v6.318A17.278 17.278 0 0 1 24 43.5h0A17.282 17.282 0 0 1 8.215 33.258"/>`),
		g.Group(children),
	)
}