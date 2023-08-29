package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rewind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M8.637 9.318a1 1 0 0 0 0 1.364l4.632 4.963c.62.664 1.731.226 1.731-.682V5.037c0-.908-1.112-1.346-1.731-.682L8.637 9.318Z"/><path fill-rule="evenodd" d="M10.736 10L13 7.574v4.852L10.736 10Zm-2.1.682a1 1 0 0 1 0-1.364l4.633-4.963c.62-.664 1.731-.226 1.731.682v9.926c0 .908-1.112 1.346-1.731.682l-4.632-4.963Z" clip-rule="evenodd"/><path d="M3.637 9.318a1 1 0 0 0 0 1.364l4.632 4.963c.62.664 1.731.226 1.731-.682V5.037c0-.908-1.112-1.346-1.731-.682L3.637 9.318Z"/><path fill-rule="evenodd" d="M5.736 10L8 7.574v4.852L5.736 10Zm-2.1.682a1 1 0 0 1 0-1.364L8.27 4.355c.619-.665 1.73-.226 1.73.682v9.926c0 .908-1.112 1.346-1.731.682l-4.632-4.963Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}