package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M16 32c8.822 0 16-7.178 16-16S24.822 0 16 0S0 7.178 0 16s7.178 16 16 16zm0-31c8.271 0 15 6.729 15 15s-6.729 15-15 15S1 24.271 1 16S7.729 1 16 1z"/><path d="M20.061 21.768a.498.498 0 0 0 .708 0a.5.5 0 0 0 0-.707L16 16.293V9.319a.5.5 0 0 0-1 0V16.5c0 .133.053.26.146.354l4.915 4.914z"/><circle cx="4" cy="16" r="1"/><circle cx="28" cy="16" r="1"/><circle cx="16" cy="4" r="1"/><circle cx="16" cy="28" r="1"/><circle cx="8" cy="8" r="1"/><circle cx="24" cy="24" r="1"/><circle cx="25" cy="8" r="1"/><circle cx="8" cy="24" r="1"/></g>`),
		g.Group(children),
	)
}