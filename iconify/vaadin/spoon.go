package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10.5 4.8c0-1.8-.9-4.8-3-4.8s-3 3-3 4.8c0 1.5.8 2.8 2.2 3.1c-.5 1.6-.7 4.6-.7 4.6v2c0 .8.7 1.5 1.5 1.5S9 15.3 9 14.5v-2c0-.6-.2-3.2-.7-4.6c1.4-.3 2.2-1.6 2.2-3.1z"/>`),
		g.Group(children),
	)
}