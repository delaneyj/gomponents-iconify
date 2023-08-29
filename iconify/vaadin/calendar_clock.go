package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 0h1v3H3V0zm8 0h1v3h-1V0z"/><path fill="currentColor" d="M6.6 14H1V6h13v.6c.4.2.7.4 1 .7V1h-2v3h-3V1H5v3H2V1H0v14h7.3c-.3-.3-.5-.6-.7-1z"/><path fill="currentColor" d="M14 12h-3V9h1v2h2z"/><path fill="currentColor" d="M11.5 8c1.9 0 3.5 1.6 3.5 3.5S13.4 15 11.5 15S8 13.4 8 11.5S9.6 8 11.5 8zm0-1C9 7 7 9 7 11.5S9 16 11.5 16s4.5-2 4.5-4.5S14 7 11.5 7z"/>`),
		g.Group(children),
	)
}