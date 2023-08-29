package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alignleftedge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.193 960h-640q-27 0-45.5-18.5t-18.5-45.5V640q0-26 18.5-45t45.5-19h640q26 0 45 19t19 45v256q0 27-19 45.5t-45 18.5zm-256-512h-384q-27 0-45.5-18.5t-18.5-45.5V128q0-26 18.5-45t45.5-19h384q26 0 45 19t19 45v256q0 27-19 45.5t-45 18.5zm-64-224q0-13-9.5-22.5t-22.5-9.5h-192q-13 0-22.5 9.5t-9.5 22.5v64q0 13 9.5 22.5t22.5 9.5h192q13 0 22.5-9.5t9.5-22.5v-64zm-576 800q-27 0-45.5-19t-18.5-45V64q0-26 18.5-45t45.5-19t45.5 19t18.5 45v896q0 27-18.5 45.5t-45.5 18.5z"/>`),
		g.Group(children),
	)
}