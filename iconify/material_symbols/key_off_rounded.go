package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.075 21.9L12.2 15.025q-.8 1.35-2.175 2.163T7 18q-2.5 0-4.25-1.75T1 12q0-1.65.813-3.025T3.974 6.8L2.1 4.925q-.275-.275-.275-.688t.275-.712q.3-.3.713-.3t.712.3L20.475 20.5q.275.275.287.688t-.287.712q-.275.275-.7.275t-.7-.275Zm3.5-9.925q0 .2-.063.375t-.212.325l-2.6 2.6q-.15.15-.312.225t-.388.075q-.225 0-.387-.063T18.3 15.3L17 14l-.075.1l-4.1-4.1H20.6q.2 0 .388.075t.312.2l.975.975q.15.15.225.337t.075.388ZM7 15q1.075 0 1.875-.663T9.9 12.726l-.563-.563l-1.25-1.25l-1.25-1.25l-.562-.562q-1.05.225-1.663 1.063T4 12q0 1.25.875 2.125T7 15Z"/>`),
		g.Group(children),
	)
}