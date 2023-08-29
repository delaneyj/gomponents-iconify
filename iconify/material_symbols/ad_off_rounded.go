package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5.825L2.075 4.9Q1.8 4.625 1.8 4.212t.3-.712q.275-.275.7-.275t.7.275l17 17q.3.3.288.7t-.313.7q-.3.275-.7.288t-.7-.288l-.9-.9H5Zm0-2h11.175L5 7.825V19Zm16-.825l-2-2V8h-8.175l-5-5H19q.825 0 1.413.588T21 5v13.175Z"/>`),
		g.Group(children),
	)
}