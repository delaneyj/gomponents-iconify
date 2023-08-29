package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Harddrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 1H3L.3 9h15.4zM0 10v5h16v-5H0zm3 3H2v-1h1v1zm4 0H4v-1h3v1z"/>`),
		g.Group(children),
	)
}