package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Roublesquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-288-896h-224q-26 0-45 19t-19 45v256h-64q-26 0-45 19t-19 45.5t19 45t45 18.5h64v64h-64q-26 0-45 19t-19 45.5t19 45t45 18.5h64v64q0 27 18.5 45.5t45 18.5t45.5-18.5t19-45.5v-64h128q26 0 45-18.5t19-45t-19-45.5t-45-19h-128v-64h160q93 0 158.5-65.5t65.5-158.5t-65.5-158.5t-158.5-65.5zm0 320h-160V256h160q40 0 68 28t28 68t-28 68t-68 28z"/>`),
		g.Group(children),
	)
}