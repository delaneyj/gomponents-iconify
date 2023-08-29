package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StylusNoteRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15.275Q3 16 3.5 16.4t1.65.525q.4.05.638.362T6 18q-.025.425-.3.7t-.675.225Q3 18.675 2 17.762t-1-2.487q0-1.625 1.338-2.638t3.712-1.212q.975-.075 1.463-.313T8 10.45q0-.55-.525-.862T5.75 9.1q-.4-.05-.638-.375T4.925 8q.05-.425.35-.688t.7-.212q2.075.3 3.05 1.113T10 10.45q0 1.325-.963 2.075t-2.837.9q-1.6.125-2.4.588T3 15.275Zm10.875 2.975L9.75 14.125L18.375 5.5q.5-.5 1.188-.5t1.187.5l1.75 1.75q.5.5.5 1.188t-.5 1.187l-8.625 8.625ZM8.975 20q-.425.1-.75-.225T8 19.025l.775-3.775l3.95 3.95l-3.75.8Z"/>`),
		g.Group(children),
	)
}