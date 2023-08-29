package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Socket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.969 12.979a4.972 4.972 0 0 1-4.967-4.968A4.975 4.975 0 0 1 8.969 3.04a4.977 4.977 0 0 1 4.969 4.971c0 2.74-2.23 4.968-4.969 4.968zm.039-9.076c-2.247 0-4.075 1.846-4.075 4.114s1.828 4.113 4.075 4.113c2.244 0 4.074-1.846 4.074-4.113c0-2.268-1.83-4.114-4.074-4.114z"/><path d="M16.916 15.918H1V0h15.916v15.918zm-14.947-.887h14.062V.937H1.969v14.094z"/><path d="M7 7h.969v1.812H7zm3 0h.969v1.812H10z"/></g>`),
		g.Group(children),
	)
}