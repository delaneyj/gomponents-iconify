package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.356 896v64q0 26-19 45t-45.5 19t-45-19t-18.5-45v-64h-640v64q0 26-19 45t-45.5 19t-45-19t-18.5-45v-64q-26 0-45-19t-19-45V64q0-27 19-45.5t45-18.5h896q26 0 45 18.5t19 45.5v768q0 26-18.5 45t-45.5 19zm-768-128h64q27 0 45.5-19t18.5-45.5t-18.5-45t-45.5-18.5h-64q-26 0-45 18.5t-19 45t19 45.5t45 19zm704-576q0-27-19-45.5t-45-18.5h-640q-27 0-45.5 18.5t-18.5 45.5v256q0 26 18.5 45t45.5 19h640q26 0 45-19t19-45V192zm-64 448h-64q-26 0-45 18.5t-19 45t19 45.5t45 19h64q27 0 45.5-19t18.5-45.5t-18.5-45t-45.5-18.5z"/>`),
		g.Group(children),
	)
}