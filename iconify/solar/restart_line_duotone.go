package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestartLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" clip-path="url(#solarRestartLineDuotone0)"><path d="M19.729 10.929a8 8 0 1 1-2.072-3.585l.707.706" opacity=".5"/><path d="M14.121 8.05h4.243V3.808"/></g><defs><clipPath id="solarRestartLineDuotone0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}