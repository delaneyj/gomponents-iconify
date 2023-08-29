package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileFocus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2ZM30 4l10 10"/><path d="m24 19l2.523 5.527l6.037.692l-4.477 4.108l1.207 5.954L24 32.293l-5.29 2.988l1.207-5.954l-4.477-4.108l6.037-.692L24 19Z"/></g>`),
		g.Group(children),
	)
}