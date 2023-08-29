package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileFocusOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 23v-9L31 4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h12"/><path d="m34 27l2.523 5.527l6.036.692l-4.476 4.108l1.207 5.954L34 40.293l-5.29 2.988l1.207-5.954l-4.477-4.108l6.037-.692L34 27ZM30 4v10h10"/></g>`),
		g.Group(children),
	)
}