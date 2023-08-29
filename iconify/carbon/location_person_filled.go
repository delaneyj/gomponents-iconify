package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationPersonFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M20 19h-2v-2h-4v2h-2v-2a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2Zm-1-8a3 3 0 1 0-3 3a3.003 3.003 0 0 0 3-3Zm-2 0a1 1 0 1 1-1-1a1.001 1.001 0 0 1 1 1Z"/><circle cx="16" cy="11" r="1" fill="currentColor"/><path fill="currentColor" d="M16 2A11.013 11.013 0 0 0 5 13a10.889 10.889 0 0 0 2.216 6.6s.3.395.349.452L16 30l8.439-9.953c.044-.053.345-.447.345-.447l.001-.003A10.885 10.885 0 0 0 27 13A11.013 11.013 0 0 0 16 2Zm0 6a3 3 0 1 1-3 3a3.003 3.003 0 0 1 3-3Zm4 11h-2v-2h-4v2h-2v-2a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}