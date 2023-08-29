package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cookie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.152 4.087c0-.367-.052-.733-.152-1.087c4.968.005 8.994 4.04 9 9c.016 4.962-4.03 8.983-9 9c-4.97.016-8.984-4.037-9-9c1.112.236 2.27-.002 3.15-.72a3.831 3.831 0 0 0 1.42-3.006a4.5 4.5 0 0 0-.07-.781a3.277 3.277 0 0 0 3.07-.351a3.69 3.69 0 0 0 1.582-3.055Zm-9.15 2.915V7H3v.002h.002Zm5-4V3H8v.002h.002Zm-4 0V3H4v.002h.002Z"/><path d="M10.002 17.002V17H10v.002h.002Zm5-2V15H15v.002h.002Zm-4-3V12H11v.002h.002Zm5-2V10H16v.002h.002Zm-13-3V7H3v.002h.002Zm5-4V3H8v.002h.002Zm-4 0V3H4v.002h.002Z"/></g>`),
		g.Group(children),
	)
}