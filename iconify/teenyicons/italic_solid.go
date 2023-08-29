package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ItalicSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.91 2H4V1h9v1H8.924L7.09 13H11v1H2v-1h4.076L7.91 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}