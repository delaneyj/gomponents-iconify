package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ambience(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.76 4.5v37.44m3.12-24.96V43.5M24 8.4v26.52m3.12-23.4v26.52M30.24 8.4v26.52m3.12-16.38v11.7M14.64 12.3v24.18M11.52 20.1v7.8m24.96-7.8v7.8"/>`),
		g.Group(children),
	)
}