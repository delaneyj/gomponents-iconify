package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><circle cx="24" cy="24" r="20" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 38L35 35"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 10L13 13"/><path fill="#2F88FF" d="M21.1429 28L18 17L14.8571 28H21.1429Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 31L14.8571 28M22 31L21.1429 28M21.1429 28L18 17L14.8571 28M21.1429 28H14.8571"/><path fill="#2F88FF" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M35 24C35 29 31.4183 31 27 31V17C31.4183 17 35 19 35 24Z"/></g>`),
		g.Group(children),
	)
}