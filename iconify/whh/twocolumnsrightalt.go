package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twocolumnsrightalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.488 1024h-384q-53 0-90.5-37.5t-37.5-90.5V128q0-53 37.5-90.5t90.5-37.5h384q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm0-864q0-13-9.5-22.5t-22.5-9.5h-320q-13 0-22.5 9.5t-9.5 22.5v704q0 13 9.5 22.5t22.5 9.5h320q13 0 22.5-9.5t9.5-22.5V160zm-704 864h-128q-27 0-45.5-19t-18.5-45V832q0-27 18.5-45.5t45.5-18.5h128q26 0 45 18.5t19 45.5v128q0 26-19 45t-45 19zm0-384h-128q-27 0-45.5-18.5T.488 576V448q0-27 18.5-45.5t45.5-18.5h128q26 0 45 18.5t19 45.5v128q0 26-19 45t-45 19zm0-384h-128q-27 0-45.5-19t-18.5-45V64q0-27 18.5-45.5T64.488 0h128q26 0 45 18.5t19 45.5v128q0 26-19 45t-45 19z"/>`),
		g.Group(children),
	)
}