package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.654 257h-896q-27 0-45.5-19t-18.5-45.5t18.5-45t45.5-18.5l448-129l448 129q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19zm-832 576q-27 0-45.5-19t-18.5-45V385q0-27 18.5-45.5t45.5-18.5h128q27 0 45.5 18.5t18.5 45.5v384q0 26-19 45t-45 19h-128zm320 0q-27 0-45.5-19t-18.5-45V385q0-27 19-45.5t45-18.5h128q27 0 45.5 18.5t18.5 45.5v384q0 26-18.5 45t-45.5 19h-128zm320 0q-26 0-45-19t-19-45V385q0-27 19-45.5t45-18.5h128q27 0 45.5 18.5t18.5 45.5v384q0 26-18.5 45t-45.5 19h-128zm-704 64h896q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19h-896q-27 0-45.5-19t-18.5-45.5t18.5-45t45.5-18.5z"/>`),
		g.Group(children),
	)
}