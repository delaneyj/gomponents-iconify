package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0h4v3H0V0zm0 4h4v3H0V4zm0 8h4v3H0v-3zm0-4h4v3H0V8zm5-8h11v3H5V0zm0 4h11v3H5V4zm0 8h11v3H5v-3zm0-4h11v3H5V8z"/>`),
		g.Group(children),
	)
}