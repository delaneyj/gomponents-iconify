package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Woman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.3 12.535a.7.7 0 1 1 0 1.399l-6.257-.001V19.3a.7.7 0 0 1-1.4 0v-5.367l-6.943.001a.7.7 0 1 1 0-1.4h14.6ZM10.343 0c3.148 0 5.7 2.55 5.7 5.696a5.698 5.698 0 0 1-5.7 5.697a5.698 5.698 0 1 1 0-11.393Zm0 1.4a4.298 4.298 0 1 0 4.3 4.296a4.298 4.298 0 0 0-4.3-4.297Z"/>`),
		g.Group(children),
	)
}