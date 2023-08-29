package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vuetify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.094 0L12 11.596L16.906 0H7.094zM1.5 3.5L12 24L22.5 3.5H17L12 15L7 3.5z"/>`),
		g.Group(children),
	)
}