package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanToolOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.05 23q-.75 0-1.4-.338T7.575 21.7l-5.9-8.65q-.2-.3-.175-.65t.3-.6q.475-.475 1.125-.55t1.175.3L7 13.575V4q0-.425.288-.713T8 3q.425 0 .713.288T9 4v11.5q0 .6-.537.888t-1.038-.063l-2.125-1.5l3.925 5.725q.125.2.35.325t.475.125H17q.825 0 1.413-.588T19 19V5q0-.425.288-.713T20 4q.425 0 .713.288T21 5v14q0 1.65-1.175 2.825T17 23h-6.95ZM12 1q.425 0 .713.288T13 2v9q0 .425-.288.713T12 12q-.425 0-.713-.288T11 11V2q0-.425.288-.713T12 1Zm4 1q.425 0 .713.288T17 3v8q0 .425-.288.713T16 12q-.425 0-.713-.288T15 11V3q0-.425.288-.713T16 2Zm-3.85 14.5Z"/>`),
		g.Group(children),
	)
}