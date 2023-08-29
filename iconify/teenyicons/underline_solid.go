package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnderlineSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 7.5V1h1v6.5a3.5 3.5 0 1 0 7 0V1h1v6.5a4.5 4.5 0 0 1-9 0ZM13 13v1H2v-1h11Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}