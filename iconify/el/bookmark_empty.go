package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M234.364 0v1200L600 834.364L965.636 1200V0H234.364zm75 75h581.271v937.5L600 721.864L309.364 1012.5V75z"/>`),
		g.Group(children),
	)
}