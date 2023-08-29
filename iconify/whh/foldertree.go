package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foldertree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M639.998 736q0-13 9.5-22.5t22.5-9.5h128q13 0 22.5 9.5t9.5 22.5v32h160q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5h-320q-13 0-22.5-9.5t-9.5-22.5v-32h-416q-13 0-22.5-9.5t-9.5-22.5V384h-160q-13 0-22.5-9.5t-9.5-22.5V32q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v32h160q13 0 22.5 9.5t9.5 22.5v256q0 13-9.5 22.5t-22.5 9.5h-160v128h384V352q0-13 9.5-22.5t22.5-9.5h128q13 0 22.5 9.5t9.5 22.5v32h160q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5h-320q-13 0-22.5-9.5t-9.5-22.5v-32h-384v320h384V736z"/>`),
		g.Group(children),
	)
}