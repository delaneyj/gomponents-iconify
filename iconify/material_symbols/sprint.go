package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sprint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.4 20L4 18.6L13.6 9H11v2H9V7h5.825q.4 0 .775.15t.65.425l3 2.975q.675.675 1.65 1.05t2.1.4v2q-1.55 0-2.813-.475T17.95 12.1l-1-1.05l-2.2 2.2L17 15.5l-6.55 3.775l-1-1.725l4.3-2.475l-1.7-1.7L5.4 20ZM3 13v-2h5v2H3Zm-2-3V8h5v2H1Zm18.475-2q-.825 0-1.425-.588T17.45 6q0-.825.6-1.413T19.475 4q.825 0 1.425.588T21.5 6q0 .825-.6 1.413T19.475 8ZM3 7V5h5v2H3Z"/>`),
		g.Group(children),
	)
}