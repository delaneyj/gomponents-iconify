package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarouselHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 26H10a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2zM10 8v16h12V8zM4 24H0v-2h4V10H0V8h4a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2zm28 0h-4a2 2 0 0 1-2-2V10a2 2 0 0 1 2-2h4v2h-4v12h4z"/>`),
		g.Group(children),
	)
}