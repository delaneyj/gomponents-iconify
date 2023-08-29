package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopFastForwardCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M14.363 12.318a1 1 0 0 1 0 1.364l-4.632 4.963c-.62.664-1.731.226-1.731-.682V8.037c0-.908 1.112-1.346 1.731-.682l4.632 4.963Z"/><path fill-rule="evenodd" d="M12.264 13L10 10.574v4.852L12.264 13Zm2.1.682a1 1 0 0 0 0-1.364L9.73 7.355C9.111 6.69 8 7.129 8 8.037v9.926c0 .908 1.112 1.346 1.731.682l4.632-4.963Z" clip-rule="evenodd"/><path d="M19.363 12.318a1 1 0 0 1 0 1.364l-4.632 4.963c-.62.664-1.731.226-1.731-.682V8.037c0-.908 1.112-1.346 1.731-.682l4.632 4.963Z"/><path fill-rule="evenodd" d="M17.264 13L15 10.574v4.852L17.264 13Zm2.1.682a1 1 0 0 0 0-1.364L14.73 7.355c-.619-.665-1.73-.226-1.73.682v9.926c0 .908 1.112 1.346 1.731.682l4.632-4.963Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopFastForwardCircleFilled0)"/></g>`),
		g.Group(children),
	)
}