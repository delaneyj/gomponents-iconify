package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M7.293 1.393a1 1 0 0 1 1.414 0l3.536 3.536a1 1 0 0 1 0 1.414l-4.95 4.95a1 1 0 1 0 1.414 1.414l4.95-4.95a1 1 0 0 1 1.414 0l3.536 3.536a1 1 0 0 1 0 1.414l-4.95 4.95A8 8 0 1 1 2.343 6.343l4.95-4.95ZM8 3.515L3.757 7.757a6 6 0 1 0 8.486 8.486L16.485 12l-2.121-2.121l-4.243 4.242A3 3 0 1 1 5.88 9.88l4.242-4.243L8 3.515Z"/><path d="m7.293 8.464l-2.121-2.12l1.414-1.415L8.707 7.05L7.293 8.464Zm6.364 6.364l-2.121-2.12l1.414-1.415l2.121 2.121l-1.414 1.414Zm.972-8.4a1 1 0 0 1-.557-1.3l1.35-3.375l1.558.52l.819-1.706a1 1 0 0 1 1.803.866L18.02 4.727l-1.442-.48l-.65 1.624a1 1 0 0 1-1.3.557Z"/></g><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}