package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BitcoinSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 1V0h1v1h3V0h1v1h.5a3.5 3.5 0 0 1 2.21 6.215A3.501 3.501 0 0 1 9.5 14H8v1H7v-1H4v1H3v-1H2v-1h1V8H2V7h1V2H2V1h1Zm1 1v5h4.5a2.5 2.5 0 0 0 0-5H4Zm0 6h5.5a2.5 2.5 0 0 1 0 5H4V8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}