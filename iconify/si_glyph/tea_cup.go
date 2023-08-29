package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TeaCup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 2h11.968v1.976H0zm9.469 11.935H2.54C.06 13.935.02 11.5.02 11.5V5.042h11.968v6.347s-.039 2.546-2.519 2.546zm3.418-9.9h-.877v1.902c.678 0 2.095-.375 2.095 1.161v.778c0 1.149-.729 1.206-2.095 1.206v1.859h.877c1.709 0 3.092-1.091 3.092-2.436V6.468c0-1.343-1.383-2.433-3.092-2.433z"/>`),
		g.Group(children),
	)
}