package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileClinicOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M17 18v-3h2v3h3v2h-3v3h-2v-3h-3v-2h3Z"/><path fill-rule="evenodd" d="M4 13a3 3 0 0 1 3-3h22a3 3 0 0 1 3 3v6h3.718a3 3 0 0 1 2.035.796l5.282 4.875A3 3 0 0 1 44 26.876V35h-3.126a4.002 4.002 0 0 1-7.748 0h-9.252a4.002 4.002 0 0 1-7.748 0h-1.252a4.002 4.002 0 0 1-7.748 0H4V13Zm33 17a4.002 4.002 0 0 0-3.874 3H32v-5h10v5h-1.126A4.002 4.002 0 0 0 37 30Zm-.604-8.735L41.526 26H32v-5h3.718a1 1 0 0 1 .678.265ZM7 12a1 1 0 0 0-1 1v13h24V13a1 1 0 0 0-1-1H7Zm13 18a4.002 4.002 0 0 1 3.874 3H30v-5H6v5h1.126a4.002 4.002 0 0 1 7.748 0h1.252c.444-1.725 2.01-3 3.874-3Zm-9 6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm11-2a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm15 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}