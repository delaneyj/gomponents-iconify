package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListUl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 1h3v3H0V1zm0 5h3v3H0V6zm0 5h3v3H0v-3zM5 1h11v3H5V1zm0 5h11v3H5V6zm0 5h11v3H5v-3z"/>`),
		g.Group(children),
	)
}