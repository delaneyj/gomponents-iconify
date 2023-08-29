package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.49 7.52a.19.19 0 0 1 0-.08a.17.17 0 0 1 0-.07v-.09l-.06-.15a.48.48 0 0 0-.09-.11l-.09-.08h-.05l-3.94-2.49l-3.72-2.3a.85.85 0 0 0-.29-.15h-.08a.82.82 0 0 0-.27 0h-.1a1.13 1.13 0 0 0-.33.13L4 6.78l-.09.07l-.09.08l-.1.07l-.05.06l-.06.15v.15a.69.69 0 0 0 0 .2v8.73a1 1 0 0 0 .47.85l7.5 4.64l.15.06h.08a.86.86 0 0 0 .52 0h.08l.15-.06L20 17.21a1 1 0 0 0 .47-.85V7.63s.02-.07.02-.11ZM12 4.17l1.78 1.1l-5.59 3.46l-1.79-1.1Zm-1 15l-5.5-3.36V9.42l5.5 3.4Zm1-8.11l-1.91-1.15l5.59-3.47l1.92 1.19Zm6.5 4.72L13 19.2v-6.38l5.5-3.4Z"/>`),
		g.Group(children),
	)
}