package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0a8 8 0 0 1 8 8v11a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V8a8 8 0 0 1 8-8Zm6.6 10.245H3.4V18.6h13.2v-8.355Zm-6.239 1.378a1.5 1.5 0 0 1 .695 2.83v2.169a.7.7 0 0 1-1.4 0v-2.175a1.5 1.5 0 0 1 .705-2.824ZM10 1.4A6.6 6.6 0 0 0 3.4 8v.845h13.2V8A6.6 6.6 0 0 0 10 1.4Z"/>`),
		g.Group(children),
	)
}