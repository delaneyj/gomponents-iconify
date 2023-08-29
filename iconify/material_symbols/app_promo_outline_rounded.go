package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppPromoOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 11.15V8q0-.425.288-.713T12 7q.425 0 .713.288T13 8v3.15l.9-.875q.275-.275.688-.275t.712.3q.275.275.275.7t-.275.7l-2.6 2.6q-.3.3-.7.3t-.7-.3l-2.6-2.6q-.275-.275-.287-.687T8.7 10.3q.275-.275.688-.288t.712.263l.9.875ZM7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Zm0-5v3h10v-3H7Zm0-2h10V6H7v10ZM7 4h10V3H7v1Zm0 14v3v-3ZM7 4V3v1Zm3.5 16h3q.2 0 .35-.15t.15-.35q0-.2-.15-.35T13.5 19h-3q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15Z"/>`),
		g.Group(children),
	)
}