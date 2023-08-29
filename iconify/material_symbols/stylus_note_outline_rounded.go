package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StylusNoteOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.475 16.825L20.85 8.45l-1.3-1.3l-8.375 8.375l1.3 1.3ZM3 15.275Q3 16 3.5 16.4t1.65.525q.4.05.638.362T6 18q-.025.425-.3.7t-.675.225Q3 18.675 2 17.762t-1-2.487q0-1.625 1.338-2.638t3.712-1.212q.975-.075 1.463-.313T8 10.45q0-.55-.525-.862T5.75 9.1q-.4-.05-.638-.375T4.925 8q.05-.425.35-.688t.7-.212q2.075.3 3.05 1.113T10 10.45q0 1.325-.963 2.075t-2.837.9q-1.6.125-2.4.588T3 15.275Zm9.95 3.9L8.825 15.05l9.55-9.55q.5-.5 1.188-.5t1.187.5l1.75 1.75q.5.5.5 1.188t-.5 1.187l-9.55 9.55ZM8.975 20q-.425.1-.75-.225T8 19.025l.825-3.975l4.125 4.125L8.975 20Z"/>`),
		g.Group(children),
	)
}