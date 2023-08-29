package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.916 9.941V7.937H10v-1h3.029v.969h1.033v2.031l2.947.014c-.14-1.43-1.062-2.539-2.189-2.539a1.74 1.74 0 0 0-.714.168c-.694-1.04-1.831-1.72-3.118-1.72c-.207 0-.408.022-.605.056C9.701 4.783 8.514 4.03 7.158 4.03c-1.91 0-3.49 1.497-3.779 3.451c-.018 0-.034-.006-.051-.006c-1.281 0-2.318 1.108-2.318 2.476l7.906-.01z"/><path d="M12.969 9h-1.031v-.994h-.876V9H10.01v.959h1.052v.989h.876v-.989h1.031V9z"/></g>`),
		g.Group(children),
	)
}