package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.59 20.41L20.17 24l-3.59 3.59L18 29l5-5l-5-5l-1.41 1.41zm7 0L27.17 24l-3.59 3.59L25 29l5-5l-5-5l-1.41 1.41z"/><path fill="currentColor" d="M14 23H4V7.91l11.43 7.91a1 1 0 0 0 1.14 0L28 7.91V17h2V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h10ZM25.8 7L16 13.78L6.2 7Z"/>`),
		g.Group(children),
	)
}