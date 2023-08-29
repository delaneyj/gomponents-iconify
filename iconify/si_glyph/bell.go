package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.141 1.318c0-.711-.559-1.286-1.248-1.286S7.646.607 7.646 1.318c0 .08.009.16.022.237C5.002 2.594 3 4.772 3 10h12c.001-5.229-2.216-7.405-4.883-8.445c.015-.077.024-.157.024-.237zM16 13H2l.906-2h11.906L16 13zm-7.039 3a1.959 1.959 0 0 0 1.961-1.957H7A1.96 1.96 0 0 0 8.961 16z"/>`),
		g.Group(children),
	)
}