package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Facebookalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-768-192.5q0 26.5 19 45.5t45 19h256V768h-256q-26 0-45 18.5t-19 45zm704-340.5q0-18-18.5-30.5t-45.5-12.5h-128V320q0-27 19-45.5t45-18.5h64q27 0 45.5-12.5t18.5-30.5v-42q0-18-18.5-30.5t-45.5-12.5h-64q-106 0-149 43t-43 149v128h-64q-26 0-45 12.5t-19 30.5v42q0 18 19 30.5t45 12.5h64v256q0 26 12.5 45t30.5 19h42q18 0 30.5-19t12.5-45V576h128q27 0 45.5-12.5t18.5-30.5v-42zm0 277h-128v128h128q27 0 45.5-19t18.5-45.5t-19-45t-45-18.5z"/>`),
		g.Group(children),
	)
}