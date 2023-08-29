package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M10.293 4.393a1 1 0 0 1 1.414 0l3.536 3.536a1 1 0 0 1 0 1.414l-4.95 4.95a1 1 0 1 0 1.414 1.414l4.95-4.95a1 1 0 0 1 1.414 0l3.536 3.536a1 1 0 0 1 0 1.414l-4.95 4.95A8 8 0 1 1 5.343 9.343l4.95-4.95ZM11 6.515l-4.243 4.242a6 6 0 1 0 8.486 8.486L19.485 15l-2.121-2.121l-4.243 4.242A3 3 0 1 1 8.88 12.88l4.242-4.243L11 6.515Z"/><path d="m10.293 11.464l-2.121-2.12l1.414-1.415l2.121 2.121l-1.414 1.414Zm6.364 6.364l-2.121-2.12l1.414-1.415l2.121 2.121l-1.414 1.414Zm.972-8.4a1 1 0 0 1-.558-1.3l1.35-3.375l1.559.52l.819-1.706a1 1 0 0 1 1.803.866L21.02 7.727l-1.442-.48l-.65 1.624a1 1 0 0 1-1.3.557Z"/></g><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}