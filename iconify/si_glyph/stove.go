package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 0v3h14.958V0H1zm2 2H2V1h1v1zm2 0H4V1h1v1zM1 16h14.958V4.042H1V16zM5 6h7v1H5V6zM4 7.958h9v6H4v-6z"/>`),
		g.Group(children),
	)
}