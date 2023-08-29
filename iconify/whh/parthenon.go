package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parthenon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 897q13 0 22.5 9.5t9.5 22.5v64q0 13-9.5 22.5T992 1025H32q-13 0-22.5-9.5T0 993v-64q0-13 9.5-22.5T32 897h32V385H32q-13 0-22.5-9.5T0 353v-64q0-13 9.5-22.5T32 257h96L490 6q22-11 44 0l362 251h96q13 0 22.5 9.5t9.5 22.5v64q0 13-9.5 22.5T992 385h-32v512h32zM256 385h-64v512h64V385zm192 0h-64v512h64V385zm128 512h64V385h-64v512zm256-512h-64v512h64V385z"/>`),
		g.Group(children),
	)
}