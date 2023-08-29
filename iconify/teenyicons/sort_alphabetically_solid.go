package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAlphabeticallySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 13.293V0h1v13.293l2.146-2.147l.708.708l-3 3a.5.5 0 0 1-.708 0l-3-3l.708-.708L3 13.293ZM11.5 1A1.5 1.5 0 0 0 10 2.5V4h3V2.5A1.5 1.5 0 0 0 11.5 1ZM13 5v2h1V2.5a2.5 2.5 0 0 0-5 0V7h1V5h3ZM9 8h3a2 2 0 0 1 1.323 3.5A2 2 0 0 1 12 15H9V8Zm3 3a1 1 0 1 0 0-2h-2v2h2Zm-2 1h2a1 1 0 1 1 0 2h-2v-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}