package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOffF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 6.757L14.243 11L5.05 20.192a2 2 0 0 1-2.828 0L.808 18.778a2 2 0 0 1 0-2.828L10 6.757zm.904-1.924a3 3 0 0 1 .695-1.09l3.35-3.35a1 1 0 0 1 1.415 0l4.243 4.243a1 1 0 0 1 0 1.414l-3.32 3.32a3 3 0 0 1-1.134.712l-5.25-5.25zm-2.318 7.581a1 1 0 0 0 1.414 0L11.414 11A1 1 0 1 0 10 9.586L8.586 11a1 1 0 0 0 0 1.414z"/>`),
		g.Group(children),
	)
}