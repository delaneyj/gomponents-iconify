package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minecraftalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm0-960h-768q-26 0-45 18.5t-19 45.5v256h64v128h128V256h128V128h128v256h128V256h128v128h128V256h64V128q0-27-18.5-45.5t-45.5-18.5z"/>`),
		g.Group(children),
	)
}