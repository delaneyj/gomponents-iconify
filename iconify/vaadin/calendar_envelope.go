package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarEnvelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 0h1v2H3V0zm6 0h1v2H9V0z"/><path fill="currentColor" d="M13 7V1h-2v2H8V1H5v2H2V1H0v12h4v3h12V7h-3zm-9 5H1V5h11v2H4v5zm1-1.8l2.6 1.5L5 14.3v-4.1zm.7 4.8l2.8-2.8l1.5.9l1.5-.8l2.8 2.8H5.7zm9.3-.7l-2.6-2.6l2.6-1.4v4zm0-5.1l-5 2.7L5 9V8h10v1.2zm.4.4z"/>`),
		g.Group(children),
	)
}