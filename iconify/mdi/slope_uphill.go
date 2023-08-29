package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlopeUphill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 13v9H2v-3l20-6m-.32-5.94l-4.82-2.6l.84 2.78l-10.12 3a2.895 2.895 0 0 0-4.05-.62a2.895 2.895 0 0 0-.62 4.05c.94 1.3 2.76 1.57 4.05.62c.71-.51 1.14-1.32 1.18-2.2l10.12-3l.84 2.78l2.58-4.81Z"/>`),
		g.Group(children),
	)
}