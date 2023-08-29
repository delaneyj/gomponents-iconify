package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vidyoconnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.09 12.09H35.9V35.9H12.09z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.06 17.06h13.88v13.88H17.06zm6.94 0v-4.97M30.94 24h4.97M24 30.94v4.97M17.06 24h-4.97"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.84 35.91v6.59H7.5a2 2 0 0 1-2-2V25.16h6.59m0-2.32H5.5V7.5a2 2 0 0 1 2-2h15.34v6.59m2.32 0V5.5H40.5a2 2 0 0 1 2 2v15.34h-6.59m0 2.32h6.59V40.5a2 2 0 0 1-2 2H25.16v-6.59M21.5 21.5h5v5h-5z"/>`),
		g.Group(children),
	)
}