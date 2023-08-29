package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RandSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 1h5a4 4 0 0 1 2.117 7.395A3.5 3.5 0 0 1 12 11.5V14h-1v-2.5A2.5 2.5 0 0 0 8.5 9H4v5H3V1Zm1 7h4a3 3 0 0 0 0-6H4v6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}