package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m28 13l-6.5 11l-1.625 2.75l-.813 1.375M15 35l.813-1.375M12 29h11m7 0h6M19 13l3.5 5.5l.875 1.375M33 35l-3.5-5.5l-1.75-2.75l-.875-1.375"/></g>`),
		g.Group(children),
	)
}