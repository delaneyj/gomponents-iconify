package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeatPumpSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.25 16.95V13.8l-2.225 2.225q.5.35 1.063.575t1.162.35Zm1.5-.025q.6-.1 1.163-.325t1.062-.575L12.75 13.8v3.125Zm3.275-1.95q.35-.5.563-1.062t.337-1.163H13.8l2.225 2.225ZM13.8 11.25h3.125q-.125-.575-.338-1.137t-.562-1.063l-2.225 2.2Zm-1.05-1.05l2.225-2.225q-.5-.35-1.063-.575t-1.162-.35v3.15ZM12 13q.425 0 .712-.287T13 12q0-.425-.288-.713T12 11q-.425 0-.713.288T11 12q0 .425.288.713T12 13Zm-.75-2.8V7.075q-.6.1-1.163.325t-1.062.575L11.25 10.2Zm-4.175 1.05H10.2l-2.225-2.2q-.35.5-.575 1.05t-.325 1.15Zm.9 3.725L10.2 12.75H7.05q.125.6.35 1.163t.575 1.062ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}