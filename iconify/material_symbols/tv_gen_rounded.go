package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvGenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.675 19l-.55 1.65q-.05.15-.175.25t-.3.1H4.5q-.2 0-.35-.15T4 20.5V19q-.825 0-1.412-.588T2 17V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v11q0 .825-.588 1.413T20 19v1.525q0 .2-.138.338t-.337.137h-.175q-.15 0-.275-.088t-.175-.237L18.35 19H5.675Z"/>`),
		g.Group(children),
	)
}