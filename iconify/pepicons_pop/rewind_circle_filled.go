package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopRewindCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M11.637 12.318a1 1 0 0 0 0 1.364l4.632 4.963c.62.664 1.731.226 1.731-.682V8.037c0-.908-1.112-1.346-1.731-.682l-4.632 4.963Z"/><path fill-rule="evenodd" d="M13.736 13L16 10.574v4.852L13.736 13Zm-2.1.682a1 1 0 0 1 0-1.364l4.633-4.963c.62-.664 1.731-.226 1.731.682v9.926c0 .908-1.112 1.346-1.731.682l-4.632-4.963Z" clip-rule="evenodd"/><path d="M6.637 12.318a1 1 0 0 0 0 1.364l4.632 4.963c.62.664 1.731.226 1.731-.682V8.037c0-.908-1.112-1.346-1.731-.682l-4.632 4.963Z"/><path fill-rule="evenodd" d="M8.736 13L11 10.574v4.852L8.736 13Zm-2.1.682a1 1 0 0 1 0-1.364l4.633-4.963c.62-.664 1.731-.226 1.731.682v9.926c0 .908-1.112 1.346-1.731.682l-4.632-4.963Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopRewindCircleFilled0)"/></g>`),
		g.Group(children),
	)
}