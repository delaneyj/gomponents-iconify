package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HowToReg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.55 20.4l-3.45-3.45l1.4-1.4l2.05 2.05l5.05-5.05l1.4 1.4l-6.45 6.45ZM10 12q-1.65 0-2.825-1.175T6 8q0-1.65 1.175-2.825T10 4q1.65 0 2.825 1.175T14 8q0 1.65-1.175 2.825T10 12Zm2.85 1.3L9.2 16.95L12.25 20H2v-2.8q0-.825.425-1.55t1.175-1.1q1.275-.65 2.875-1.1T10 13q.75 0 1.463.075t1.387.225Z"/>`),
		g.Group(children),
	)
}