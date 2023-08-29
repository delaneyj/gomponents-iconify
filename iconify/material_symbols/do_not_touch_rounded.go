package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoNotTouchRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.05 23q-.75 0-1.4-.338T7.575 21.7l-5.9-8.65q-.2-.3-.175-.65t.3-.6q.475-.475 1.125-.55t1.175.3L7 13.575v-3.75L1.375 4.2q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3l-.4-.375q-.5.375-1.1.588T17 23h-6.95ZM21 18.15l-5.975-5.975H15V3q0-.425.288-.713T16 2q.425 0 .713.288T17 3v9.175h2V5q0-.425.288-.713T20 4q.425 0 .713.288T21 5v13.15Zm-8-8l-2-2V2q0-.425.288-.713T12 1q.425 0 .713.288T13 2v8.15Zm-4-4l-2-2V4q0-.425.288-.713T8 3q.425 0 .713.288T9 4v2.15Z"/>`),
		g.Group(children),
	)
}