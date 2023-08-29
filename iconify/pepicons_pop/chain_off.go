package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChainOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M10.392 11.564a2 2 0 0 1-.587-2.767l2.181-3.353a2 2 0 1 1 3.353 2.18l-2.18 3.354a2 2 0 0 1-2.768.586Zm-2.442.202a3.999 3.999 0 0 1 .179-4.06l2.18-3.353a4 4 0 1 1 6.707 4.362l-2.18 3.353a4 4 0 0 1-6.886-.302Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5.444 15.503a2 2 0 0 1-.587-2.767l2.181-3.353a2 2 0 1 1 3.353 2.18l-2.18 3.354a2 2 0 0 1-2.767.586Zm-2.442.202a3.999 3.999 0 0 1 .179-4.06l2.18-3.352a4 4 0 1 1 6.707 4.361l-2.18 3.353a4 4 0 0 1-6.886-.302Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}