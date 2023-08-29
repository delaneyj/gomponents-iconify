package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrpanoOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.85 15.25q1.2-.125 2.488-.188T12 15q1.35 0 2.638.063t2.512.187q.425.05.588-.25t-.113-.65l-3.25-3.9q-.15-.175-.375-.175t-.375.175L11.15 13.4l-1.625-1.95q-.125-.175-.363-.175t-.387.2l-2.4 2.875q-.275.35-.125.65t.6.25ZM3 20q-.425 0-.713-.288T2 19V5q0-.425.288-.713T3 4q.425 0 2.725.75T12 5.5q4 0 6.288-.75T21 4q.425 0 .713.288T22 5v14q0 .425-.288.713T21 20q-.425 0-2.713-.75T12 18.5q-3.975 0-6.275.75T3 20Zm1-2.35q1.95-.575 3.963-.863T12 16.5q2.025 0 4.038.288T20 17.65V6.375q-1.95.575-3.963.85T12 7.5q-2.025 0-4.038-.275T4 6.375V17.65ZM12 12Z"/>`),
		g.Group(children),
	)
}