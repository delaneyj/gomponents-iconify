package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScubaDivingRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15q-.825 0-1.413-.588T1 13q0-.825.588-1.413T3 11q.825 0 1.413.588T5 13q0 .825-.588 1.413T3 15Zm5.9-4.9q-.6.175-1.137-.138T7.05 9.05q-.175-.6.138-1.15t.912-.7l3.575-.95q.4-.125.75.1t.475.625l.275.95q.125.4-.1.75t-.625.475l-3.55.95ZM3.2 22.4q-.35-.25-.4-.65T3 21l2.25-3l.775-4.45q.075-.6.475-1.062t1.025-.613L17 9l2-4l2.5-2.5q.2-.2.488-.212t.537.237q.2.2.225.488t-.175.487L20.5 5.9l-1.3 3.975q-.1.3-.287.55t-.463.45l-4.2 2.95q-.125.1-.263.163t-.287.112l-5.55 1.75L7 19l-2.4 3.2q-.25.35-.65.4t-.75-.2Z"/>`),
		g.Group(children),
	)
}