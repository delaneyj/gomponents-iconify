package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsWNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM10.027 12.329a2 2 0 0 1 3.946-.658l3.076 18.46l5.214-9.123a2 2 0 0 1 3.473 0l5.214 9.124l3.077-18.461a2 2 0 0 1 3.946.658l-4 24a2 2 0 0 1-3.71.663L24 26.031l-6.264 10.961a2 2 0 0 1-3.709-.663l-4-24Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsWNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}