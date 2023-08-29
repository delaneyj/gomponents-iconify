package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchitectureRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.1 11.9l-2.7 7.45q-.025.075-.125.175l-.8.775q-.225.225-.512.113T6.625 20l-.1-1.075l.025-.225l2.8-7.75q.375.35.813.588t.937.362Zm1.8 0q.5-.125.938-.363t.812-.587l2.8 7.75q.025.075.025.225l-.1 1.075q-.05.3-.338.413t-.512-.113l-.8-.775l-.125-.175l-2.7-7.45ZM12 11q-1.25 0-2.125-.875T9 8q0-.975.563-1.738T11 5.2V4q0-.425.288-.713T12 3q.425 0 .713.288T13 4v1.2q.875.3 1.438 1.063T15 8q0 1.25-.875 2.125T12 11Zm0-2q.425 0 .713-.288T13 8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8q0 .425.288.713T12 9Z"/>`),
		g.Group(children),
	)
}