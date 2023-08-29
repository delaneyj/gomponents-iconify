package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindPowerOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 23q0-.825.588-1.413T12 21v-5.2q-.3-.125-.563-.288t-.487-.362L8.05 17L1 15v-4h9.8q.25-.275.55-.475T12 10.2V5.9L17.35.925l3.375 2.125l-5.125 8.475q.175.3.263.65t.112.725l3.525.85l3.55 6.35l-2.825 2.825l-6.225-6.2V21q.825 0 1.413.588T16 23h-6ZM4 5V3h6v2H4Zm10 1.775V10.2q.025 0 .038.013t.037.012L18 3.675l-.4-.25l-3.6 3.35ZM1 9V7h5v2H1Zm12 5q.425 0 .713-.288T14 13q0-.425-.288-.713T13 12q-.425 0-.713.288T12 13q0 .425.288.713T13 14Zm-5.275.85l2.325-1.4q-.025-.125-.025-.225V13H3v.5l4.725 1.35Zm10.475.625l-2.85-.65q-.05.05-.1.125t-.1.125l5.075 5.05l.35-.35l-2.375-4.3ZM3 21v-2h5v2H3Zm10-8Zm1-2.775ZM10.05 13Zm5.1 2.075Z"/>`),
		g.Group(children),
	)
}