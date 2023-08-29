package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayPassRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-.65 0-1.175.388t-.725.987q-.05.275-.262.45T9.35 22H7q-.825 0-1.412-.588T5 20V4q0-.825.588-1.413T7 2h2.35q.275 0 .488.175t.262.45q.2.6.725.988T12 4q.65 0 1.175-.388t.725-.987q.05-.275.263-.45T14.65 2H17q.825 0 1.413.587T19 4v16q0 .825-.588 1.413T17 22h-2.35q-.275 0-.488-.175t-.262-.45q-.2-.6-.725-.988T12 20Z"/>`),
		g.Group(children),
	)
}