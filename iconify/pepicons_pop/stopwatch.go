package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stopwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10 5.5a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11ZM2.5 11a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Z"/><path d="M2.793 3.793a1 1 0 0 1 1.414 0l1.5 1.5a1 1 0 0 1-1.414 1.414l-1.5-1.5a1 1 0 0 1 0-1.414Zm11.5 2.914a1 1 0 0 0 1.414 0l1.5-1.5a1 1 0 0 0-1.414-1.414l-1.5 1.5a1 1 0 0 0 0 1.414ZM13.28 8.375a1 1 0 0 1-.155 1.406l-2.5 2a1 1 0 0 1-1.25-1.562l2.5-2a1 1 0 0 1 1.406.156ZM9 4V2h2v2H9Z"/><path d="M7.5 2a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1Z"/></g>`),
		g.Group(children),
	)
}