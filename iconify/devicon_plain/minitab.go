package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minitab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M0 0v128h128V0H0zm34.6 106.5H14.8v-54h19.7v54zm27.1 0H42v-87h19.7v87zm27.1 0H69.1V37.6h19.7v68.9zm27.2 0H96.3V68.6H116v37.9z"/>`),
		g.Group(children),
	)
}