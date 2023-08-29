package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PulmonologyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-1.275 0-2.138-.863T2 18v-4.175L4.625 6.8q.3-.825 1.05-1.312T7.3 5q1.125 0 1.913.813T10 7.775V9H8V7.775q0-.325-.225-.55T7.25 7q-.25 0-.463.138T6.5 7.5L4 14.2V18q0 .425.288.713T5 19h3q.425 0 .713-.288T9 18v-2h2v2q0 1.275-.875 2.138T8 21H5Zm13.975 0h-3q-1.25 0-2.125-.863T12.975 18v-2h2v2q0 .425.288.713t.712.287h3q.425 0 .713-.288t.287-.712v-3.8l-2.5-6.7q-.1-.225-.3-.363T16.725 7q-.325 0-.537.225t-.213.55V9h-2V7.775q0-1.15.788-1.963T16.674 5q.875 0 1.613.488T19.35 6.8l2.625 7.025V18q0 1.275-.875 2.138T18.975 21ZM8 12.6Zm7.975 0ZM12 11.425L9.4 14L8 12.6l3-3V2h2v7.6l3 3l-1.425 1.4L12 11.425Z"/>`),
		g.Group(children),
	)
}