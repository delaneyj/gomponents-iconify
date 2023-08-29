package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linux(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-352q-13 0-22.5-9.5t-9.5-22.5V864q0-13 9.5-22.5t22.5-9.5h224q26 0 45-19t19-45V256q0-27-19-45.5t-45-18.5h-512q-26 0-45 18.5t-19 45.5v32q0 13-9.5 22.5t-22.5 9.5h-128q-13 0-22.5-9.5T.428 288V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-864-640h128q13 0 22.5 9.5t9.5 22.5v352q0 26 19 45t45 19h160q13 0 22.5 9.5t9.5 22.5v128q0 13-9.5 22.5t-22.5 9.5h-288q-53 0-90.5-37.5T.428 896V416q0-13 9.5-22.5t22.5-9.5z"/>`),
		g.Group(children),
	)
}