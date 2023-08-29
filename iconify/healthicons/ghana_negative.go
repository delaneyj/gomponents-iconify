package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhanaNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsGhanaNegative0)"><path d="M20.173 14.732a9.996 9.996 0 0 1 1.827-.56v19.596a9.995 9.995 0 0 1-6.315-4.242a10 10 0 0 1 4.488-14.794Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 6a2 2 0 0 1 2 2v2.114a14 14 0 0 1 7.9 3.957a2 2 0 0 1-2.829 2.828A10 10 0 0 0 26 14.172v19.596a9.999 9.999 0 0 0 5.071-2.727A2 2 0 0 1 33.9 33.87a14 14 0 0 1-7.9 3.957V40a2 2 0 1 1-4 0v-2.173a14 14 0 0 1 0-27.713V8a2 2 0 0 1 2-2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsGhanaNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}