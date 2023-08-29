package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricRickshawRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17q-.975 0-1.738-.563T3.2 15H3q-.825 0-1.413-.588T1 13V5q0-.825.588-1.413T3 3h12.05q.45 0 .85.175t.7.525l3.95 4.75q.225.275.338.588T21 9.7v1.45q.875.3 1.438 1.088T23 14q0 1.25-.875 2.125T20 17q-.975 0-1.763-.563T17.15 15h-8.3q-.35.875-1.113 1.438T6 17ZM3 8h4V5H3v3Zm6 5h5V5H9v3h2q.425 0 .713.288T12 9q0 .425-.288.713T11 10H9v3Zm7-4h2.4L16 6.1V9ZM6 15q.425 0 .713-.288T7 14q0-.425-.288-.713T6 13q-.425 0-.713.288T5 14q0 .425.288.713T6 15Zm14 0q.425 0 .713-.288T21 14q0-.425-.288-.713T20 13q-.425 0-.713.288T19 14q0 .425.288.713T20 15Zm-7.725 7.65L7.95 20.475q-.175-.1-.138-.288T8.05 20H11v-1.2q0-.275.238-.425t.487-.025l4.325 2.175q.175.1.138.288T15.95 21H13v1.2q0 .275-.238.425t-.487.025Z"/>`),
		g.Group(children),
	)
}