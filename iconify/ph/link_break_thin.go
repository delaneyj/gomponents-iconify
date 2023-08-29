package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkBreakThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M193.46 62.54a36.06 36.06 0 0 0-50.92 0l-11.65 12.22a4 4 0 0 1-5.78-5.52L136.82 57a44 44 0 1 1 62.29 62.15l-12.35 11.78a4 4 0 1 1-5.52-5.78l12.28-11.72a36 36 0 0 0-.06-50.89Zm-68.35 118.7l-11.65 12.22a36 36 0 0 1-51-50.85l12.28-11.72a4 4 0 0 0-5.52-5.78l-12.33 11.78A44 44 0 1 0 119.18 199l11.71-12.28a4 4 0 1 0-5.78-5.52ZM208 156h-24a4 4 0 0 0 0 8h24a4 4 0 0 0 0-8ZM48 100h24a4 4 0 0 0 0-8H48a4 4 0 0 0 0 8Zm112 80a4 4 0 0 0-4 4v24a4 4 0 0 0 8 0v-24a4 4 0 0 0-4-4ZM96 76a4 4 0 0 0 4-4V48a4 4 0 0 0-8 0v24a4 4 0 0 0 4 4Z"/>`),
		g.Group(children),
	)
}