package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VpnKeyOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.625L12.15 15q-.8 1.35-2.138 2.175T7 18q-2.5 0-4.25-1.75T1 12q0-1.65.8-3.025T3.95 6.8L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425ZM7 14q.825 0 1.413-.588T9 12v-.175L7.175 10H7q-.825 0-1.413.588T5 12q0 .825.588 1.413T7 14Zm13.825 4l-8-8H23v4h-2v4h-.175Z"/>`),
		g.Group(children),
	)
}