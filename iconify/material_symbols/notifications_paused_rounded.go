package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationsPausedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19q-.425 0-.713-.288T4 18q0-.425.288-.713T5 17h1v-7q0-2.075 1.25-3.688T10.5 4.2v-.7q0-.625.438-1.063T12 2q.625 0 1.063.438T13.5 3.5v.7q2 .5 3.25 2.113T18 10v7h1q.425 0 .713.288T20 18q0 .425-.288.713T19 19H5Zm7 3q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Zm1.85-6q.375 0 .638-.263t.262-.637q0-.375-.263-.637t-.637-.263H12l2.55-3.15q.1-.125.15-.275t.05-.3V9.9q0-.375-.262-.638T13.85 9h-3.7q-.375 0-.638.263T9.25 9.9q0 .375.263.638t.637.262H12l-2.55 3.15q-.1.125-.15.275t-.05.3v.575q0 .375.263.638t.637.262h3.7Z"/>`),
		g.Group(children),
	)
}