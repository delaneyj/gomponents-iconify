package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StylusNoteOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.475 16.825L20.85 8.45l-1.3-1.3l-8.375 8.375l1.3 1.3ZM5.95 19q-2.5-.125-3.725-1.05T1 15.275q0-1.625 1.338-2.638t3.712-1.212q.975-.075 1.463-.313T8 10.45q0-.65-.738-.975T4.825 9L5 7q2.575.2 3.788 1.038T10 10.45q0 1.325-.963 2.075t-2.837.9q-1.6.125-2.4.588T3 15.275q0 .875.7 1.263T6.05 17l-.1 2Zm7 .175L8.825 15.05l9.55-9.55q.5-.5 1.188-.5t1.187.5l1.75 1.75q.5.5.5 1.188t-.5 1.187l-9.55 9.55ZM8.975 20q-.425.1-.75-.225T8 19.025l.825-3.975l4.125 4.125L8.975 20Z"/>`),
		g.Group(children),
	)
}