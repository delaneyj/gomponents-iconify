package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.075 17.65l-.65-.975q-.475.225-.975.338t-1.025.112q-1.95 0-3.3-1.35t-1.35-3.3q0-1.925 1.35-3.288T16.4 7.825q1.925 0 3.263 1.363T21 12.474q0 1-.4 1.9t-1.125 1.575l.65.975l-1.05.725ZM3 16.925l3.325-8.9h1.6l3.35 8.9h-1.55l-.8-2.275H5.35l-.825 2.275H3Zm13.4-1.1q.35 0 .663-.075t.612-.2L16.625 14l1.075-.7l1 1.5q.425-.475.65-1.075t.225-1.25q0-1.375-.913-2.362T16.4 9.125q-1.35 0-2.275.988t-.925 2.362q0 1.375.925 2.363t2.275.987ZM5.8 13.35h2.65l-1.275-3.6h-.1L5.8 13.35Z"/>`),
		g.Group(children),
	)
}