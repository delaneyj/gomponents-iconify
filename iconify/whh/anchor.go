package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m992 896l-74-74q-71 94-177.5 148T512 1024t-228.5-54T106 822l-74 74q-13 0-22.5-9.5T0 864V672q0-13 9.5-22.5T32 640h192q13 0 22.5 9.5T256 672l-59 59q44 63 109 104.5T448 890V576h-64q-26 0-45-18.5t-19-45t19-45.5t45-19h64v-76q-57-20-92.5-69T320 192q0-80 56-136T512 0t136 56.5T704 192q0 62-35.5 111T576 372v76h64q27 0 45.5 19t18.5 45.5t-18.5 45T640 576h-64v314q77-13 142-54.5T827 731l-59-59q0-13 9.5-22.5T800 640h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5T992 896zM511.5 128q-26.5 0-45 18.5t-18.5 45t18.5 45.5t45 19t45.5-19t19-45t-19-45t-45.5-19z"/>`),
		g.Group(children),
	)
}