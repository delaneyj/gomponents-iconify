package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchVideoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.825 14.825q.275.275.675.275t.7-.3q.275-.275.275-.7t-.275-.7l-.4-.4h4.4l-.425.425q-.275.275-.275.675t.3.7q.275.275.7.275t.7-.275l2.1-2.1q.15-.15.212-.325t.063-.375q0-.2-.063-.375T15.3 11.3l-2.125-2.125Q12.9 8.9 12.5 8.9t-.7.3q-.275.275-.275.7t.275.7l.4.4H7.8l.425-.425Q8.5 10.3 8.5 9.9t-.3-.7q-.275-.275-.7-.275t-.7.275l-2.1 2.1q-.15.15-.212.325T4.425 12q0 .2.063.375t.212.325l2.125 2.125ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h12q.825 0 1.413.588T18 6v4.5l3.15-3.15q.225-.225.537-.113T22 7.7v8.6q0 .35-.313.463t-.537-.113L18 13.5V18q0 .825-.588 1.413T16 20H4Z"/>`),
		g.Group(children),
	)
}