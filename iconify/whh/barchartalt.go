package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barchartalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.193 1024h-64q-27 0-45.5-19t-18.5-45V64q0-26 18.5-45t45.5-19h64q26 0 45 19t19 45v896q0 26-19 45t-45 19zm-256 0h-64q-27 0-45.5-19t-18.5-45V320q0-26 18.5-45t45.5-19h64q26 0 45 19t19 45v640q0 26-19 45t-45 19zm-256 0h-64q-27 0-45.5-19t-18.5-45V576q0-27 18.5-45.5t45.5-18.5h64q26 0 45 18.5t19 45.5v384q0 26-19 45t-45 19zm-256-256q26 0 45 19t19 45.5t-19 45t-45 18.5h-64v64q0 26-19 45t-45.5 19t-45-19t-18.5-45V64q0-26 18.5-45t45.5-19h128q26 0 45 19t19 45.5t-19 45t-45 18.5h-64v128h64q26 0 45 19t19 45.5t-19 45t-45 18.5h-64v128h64q26 0 45 18.5t19 45t-19 45.5t-45 19h-64v128h64z"/>`),
		g.Group(children),
	)
}