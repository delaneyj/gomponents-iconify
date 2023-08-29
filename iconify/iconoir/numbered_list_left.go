package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberedListLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 5h12M5 7V3L3.5 4.5m2 9.5h-2l1.905-2.963a.428.428 0 0 0 .072-.323C5.42 10.456 5.216 10 4.5 10c-1 0-1 .889-1 .889s0 0 0 0v.222M4 19h.5a1 1 0 0 1 1 1v0a1 1 0 0 1-1 1h-1m0-4h2L4 19m5-7h12M9 19h12"/>`),
		g.Group(children),
	)
}