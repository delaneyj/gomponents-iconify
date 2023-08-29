package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.775 22.6l-1.6-1.6H5q-.825 0-1.413-.588T3 19V5.825L1.4 4.2l1.4-1.4l18.4 18.4l-1.425 1.4ZM5 19h11.175L5 7.825V19Zm16-.825l-2-2V8h-8.175l-5-5H19q.825 0 1.413.588T21 5v13.175Z"/>`),
		g.Group(children),
	)
}