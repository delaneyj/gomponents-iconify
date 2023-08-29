package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClassifierLanguage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 10V2h-8v8h3v4h-2v2h7v-2h-3v-4zm-6-6h4v4h-4zm-4 19.41L18.59 22L16 24.59L13.41 22L12 23.41L14.59 26L12 28.59L13.41 30L16 27.41L18.59 30L20 28.59L17.41 26L20 23.41zM20 14h-8v2h3v5h2v-5h3v-2zM7 9.86a4 4 0 1 0-2 0V14H2v2h7v-2H7zM4 6a2 2 0 1 1 2 2a2 2 0 0 1-2-2z"/>`),
		g.Group(children),
	)
}