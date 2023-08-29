package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.965 10.991V8.987l5.051-.034v2.031l2.994-.033c-.14-1.43-1.062-2.539-2.189-2.539a1.74 1.74 0 0 0-.714.168c-.694-1.04-1.831-1.72-3.118-1.72c-.207 0-.408.022-.605.056C8.702 5.783 7.515 5.03 6.159 5.03c-1.91 0-3.49 1.497-3.779 3.451c-.018 0-.034-.006-.051-.006c-1.281 0-2.318 1.108-2.318 2.476l7.954.04z"/><path d="M9 10h2.959v.99H9z"/></g>`),
		g.Group(children),
	)
}