package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopMagnetCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.293 4.393a1 1 0 0 1 1.414 0l3.536 3.536a1 1 0 0 1 0 1.414l-4.95 4.95a1 1 0 1 0 1.414 1.414l4.95-4.95a1 1 0 0 1 1.414 0l3.536 3.536a1 1 0 0 1 0 1.414l-4.95 4.95A8 8 0 1 1 5.343 9.343l4.95-4.95ZM11 6.515l-4.243 4.242a6 6 0 1 0 8.486 8.486L19.485 15l-2.121-2.121l-4.243 4.242A3 3 0 1 1 8.88 12.88l4.242-4.243L11 6.515Z"/><path d="m10.293 11.464l-2.121-2.12l1.414-1.415l2.121 2.121l-1.414 1.414Zm6.364 6.364l-2.121-2.12l1.414-1.415l2.121 2.121l-1.414 1.414Zm.972-8.4a1 1 0 0 1-.558-1.3l1.35-3.375l1.559.52l.819-1.706a1 1 0 0 1 1.803.866L21.02 7.727l-1.442-.48l-.65 1.624a1 1 0 0 1-1.3.557Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopMagnetCircleFilled0)"/></g>`),
		g.Group(children),
	)
}