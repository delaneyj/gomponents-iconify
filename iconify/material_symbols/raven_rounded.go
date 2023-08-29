package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RavenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 5q-.425 0-.713.288T5 6q0 .425.288.713T6 7q.425 0 .713-.288T7 6q0-.425-.288-.713T6 5Zm4 13l-1.275 3.075q-.15.375-.537.537t-.763.013q-.375-.15-.537-.537t-.013-.763l1.075-2.6q-2.65-.7-4.3-2.85T2 10V6q0-1.65 1.175-2.825T6 2q.55 0 1.05.187t1 .388l4.775 1.95q.325.125.313.463t-.338.462L10 6.5V8l7.85 5H10q-.725 0-1.275-.45T8.05 11.4q-.075-.35-.35-.6t-.65-.25q-.425 0-.713.288t-.287.712q0 .5.3 1.113t.625.962q.55.65 1.338 1.013T10 15h11l.775 3.925q.075.425-.188.75T20.9 20h-.35q-.275 0-.475-.138t-.325-.362L19 18h-5v3q0 .425-.288.713T13 22q-.425 0-.713-.288T12 21v-3h-2Z"/>`),
		g.Group(children),
	)
}