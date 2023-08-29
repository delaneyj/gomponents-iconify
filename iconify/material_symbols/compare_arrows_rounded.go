package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompareArrowsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.175 16H3q-.425 0-.713-.288T2 15q0-.425.288-.713T3 14h6.175L7.3 12.125q-.275-.275-.275-.688t.275-.712q.3-.3.713-.3t.712.3L12.3 14.3q.15.15.213.325t.062.375q0 .2-.062.375t-.213.325l-3.6 3.6q-.3.3-.7.288t-.7-.313q-.275-.3-.287-.7t.287-.7L9.175 16Zm5.65-6l1.875 1.875q.275.275.275.688t-.275.712q-.3.3-.713.3t-.712-.3L11.7 9.7q-.15-.15-.212-.325T11.425 9q0-.2.063-.375T11.7 8.3l3.6-3.6q.3-.3.7-.287t.7.312q.275.3.288.7t-.288.7L14.825 8H21q.425 0 .713.288T22 9q0 .425-.288.713T21 10h-6.175Z"/>`),
		g.Group(children),
	)
}