package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 24H0V0h24zM15.014 8.994v7.326c0 .198 0 .234-.127.362l-1.302 1.264v.27h6.32v-.27l-1.257-1.234a.376.376 0 0 1-.143-.364v.002v-9.07a.375.375 0 0 1 .143-.36l.001-.001l1.286-1.234v-.27h-4.456l-3.176 7.924l-3.609-7.924H4.019v.271l1.502 1.813a.627.627 0 0 1 .204.53v-.003v7.126a.823.823 0 0 1-.22.707l-1.69 2.054v.27h4.8v-.27l-1.691-2.054a.852.852 0 0 1-.233-.712v.004v-6.16l4.215 9.195h.49z"/>`),
		g.Group(children),
	)
}