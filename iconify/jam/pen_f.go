package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.406 12.732L5.367 9.693c.543-.956 1.335-1.903 2.375-2.838c1.341-1.208 3.708-3.25 7.1-6.126a1.795 1.795 0 0 1 2.53 2.53c-2.88 3.398-4.921 5.765-6.125 7.1c-.935 1.037-1.882 1.828-2.841 2.373zM3.99 11.146l2.965 2.964l-1.366 1.906l-5.276 1.795l1.771-5.3l1.906-1.365z"/>`),
		g.Group(children),
	)
}