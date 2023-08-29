package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 .7V0h-.7l-3 3H5V0H3v3H0v2h3v8h8v3h2v-3h3v-2h-3V3.7l3-3zM5 5h5.3L5 10.3V5zm6 6H5.7L11 5.7V11z"/>`),
		g.Group(children),
	)
}