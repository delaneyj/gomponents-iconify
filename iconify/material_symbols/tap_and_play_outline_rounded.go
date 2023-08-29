package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TapAndPlayOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 23v-2h1V6H7v6H5V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23h-1Zm-6.75-4.25q.675.675 1.1 1.5t.575 1.775q.05.4-.225.688T10 23q-.425 0-.725-.263t-.4-.687q-.125-.525-.388-1t-.662-.875q-.4-.4-.875-.663t-1-.387q-.425-.1-.688-.4T5 18q0-.425.288-.7t.687-.225q.95.15 1.775.575t1.5 1.1Zm3.15-2.5q1.05 1.15 1.713 2.588t.837 3.087q.05.425-.237.75T14 23q-.425 0-.725-.3t-.35-.725q-.15-1.25-.663-2.35T10.95 17.65Q10 16.575 8.713 15.9t-2.788-.85q-.4-.05-.662-.35T5 14q0-.425.288-.713t.687-.237q1.9.2 3.55 1.025T12.4 16.25ZM7 4h10V3H7v1Zm0 0V3v1ZM6 23q-.425 0-.713-.288T5 22q0-.425.288-.713T6 21q.425 0 .713.288T7 22q0 .425-.288.713T6 23Z"/>`),
		g.Group(children),
	)
}