package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberedListRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 5h12m5.5 2V3L19 4.5m2 9.5h-2l1.905-2.963a.428.428 0 0 0 .072-.323C20.92 10.456 20.716 10 20 10c-1 0-1 .889-1 .889s0 0 0 0v.222M19.5 19h.5a1 1 0 0 1 1 1v0a1 1 0 0 1-1 1h-1m0-4h2l-1.5 2M3 12h12M3 19h12"/>`),
		g.Group(children),
	)
}