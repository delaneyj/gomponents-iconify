package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HowToRegRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.55 19.975q-.2 0-.375-.063t-.325-.212l-2.05-2.05q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.35 1.35l4.35-4.35q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-5.05 5.05q-.15.15-.325.213t-.375.062ZM10 12q-1.65 0-2.825-1.175T6 8q0-1.65 1.175-2.825T10 4q1.65 0 2.825 1.175T14 8q0 1.65-1.175 2.825T10 12Zm-7 8q-.425 0-.713-.288T2 19v-1.8q0-.825.425-1.55t1.175-1.1q1.275-.65 2.875-1.1T10 13q.75 0 1.463.075t1.387.225l-2.25 2.25q-.575.575-.562 1.412t.587 1.413L12.25 20H3Z"/>`),
		g.Group(children),
	)
}