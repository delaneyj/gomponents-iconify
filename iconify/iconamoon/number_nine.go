package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.153 19.468a1 1 0 1 0 1.694 1.064l-1.694-1.064Zm6.92-7.264a1 1 0 1 0-1.693-1.064l1.694 1.064ZM8 9a4 4 0 0 1 4-4V3a6 6 0 0 0-6 6h2Zm4-4a4 4 0 0 1 4 4h2a6 6 0 0 0-6-6v2Zm4 4a4 4 0 0 1-4 4v2a6 6 0 0 0 6-6h-2Zm-4 4a4 4 0 0 1-4-4H6a6 6 0 0 0 6 6v-2Zm-.153 7.532l5.227-8.328l-1.694-1.064l-5.227 8.328l1.694 1.064Z"/>`),
		g.Group(children),
	)
}