package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanToolAltOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.95 21q-.675 0-1.287-.288T7.625 19.9L1.15 11.925l.65-.625q.475-.475 1.125-.55t1.175.3L7 13.075V3q0-.425.288-.713T8 2q.425 0 .713.288T9 3v12q0 .6-.537.888t-1.038-.063l-1.2-.85l2.95 3.65q.15.175.35.275t.425.1H17q.825 0 1.413-.588T19 17v-7q0-.425.288-.713T20 9q.425 0 .713.288T21 10v7q0 1.65-1.175 2.825T17 21H9.95ZM11 13V7q0-.425.288-.713T12 6q.425 0 .713.288T13 7v6h-2Zm4 0V8q0-.425.288-.713T16 7q.425 0 .713.288T17 8v5h-2Zm-2.375 2Z"/>`),
		g.Group(children),
	)
}