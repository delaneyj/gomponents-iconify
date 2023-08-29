package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trademark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 7h-6a1 1 0 0 0 0 2h2v7a1 1 0 0 0 2 0V9h2a1 1 0 0 0 0-2Zm11.92.62a1 1 0 0 0-.54-.54a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21L17 10.09l-2.79-2.8a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.54.54a1 1 0 0 0-.08.38v8a1 1 0 0 0 2 0v-5.59l1.79 1.8a1 1 0 0 0 1.42 0l1.79-1.8V16a1 1 0 0 0 2 0V8a1 1 0 0 0-.08-.38Z"/>`),
		g.Group(children),
	)
}