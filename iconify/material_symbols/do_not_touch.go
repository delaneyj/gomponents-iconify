package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoNotTouch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.1 2.075l19.8 19.8l-1.425 1.425l-1.1-1.1q-.5.375-1.1.588T17 23h-6.95q-.75 0-1.4-.338T7.575 21.7L1.2 12.375l.6-.575q.475-.475 1.125-.55t1.175.3L7 13.575v-3.75L.675 3.5L2.1 2.075ZM21 18.15l-6-5.975V3q0-.425.288-.713T16 2q.425 0 .713.288T17 3v9.175h2V5q0-.425.288-.713T20 4q.425 0 .713.288T21 5v13.15Zm-8-8l-2-2V2q0-.425.288-.713T12 1q.425 0 .713.288T13 2v8.15Zm-4-4l-2-2V4q0-.425.288-.713T8 3q.425 0 .713.288T9 4v2.15Z"/>`),
		g.Group(children),
	)
}