package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.797 1.272c-1.549-1.55-4.148-1.65-5.69-.107c-1.119 1.117-1.317 2.834-.812 4.233l-8.223 8.223s-.534 2.827 1.85 2.215c.058-.076 1.047-.586 1.047-.586l.023-2.289L6.215 13l.74-.875l.014-2.172h1.977l-.002-1.975l2.168-.008l.379-.258c1.439.602 3.197.162 4.321-.843c1.626-1.453 1.535-4.048-.015-5.597zm-.781 1.759h-1.047V1.984h1.047v1.047z"/>`),
		g.Group(children),
	)
}