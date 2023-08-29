package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tuneinalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 960H128q-53 0-90.5-37.5T0 832V128q0-53 37.5-90.5T128 0h704q53 0 90.5 37.5T960 128v704q0 53-37.5 90.5T832 960zM480 128q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm96 352q0-40-29-68t-71-28H128v96q0 40 29 68t71 28h348v-96zm0 160H384v192h96q40 0 68-28t28-68v-96zm256-256h-96q-40 0-68 28t-28 68v96h96q40 0 68-28t28-68v-96z"/>`),
		g.Group(children),
	)
}