package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5.903 12.646a1 1 0 0 1-.906-1.087l.472-5.185a1 1 0 0 1 1.991.181l-.471 5.186a1 1 0 0 1-1.086.905Z"/><path d="M12.646 5.903a1 1 0 0 1-.905 1.086l-5.186.471a1 1 0 0 1-.181-1.991l5.185-.472a1 1 0 0 1 1.087.906Z"/><path d="M7.172 7.172a1 1 0 0 1 1.414 0l5.657 5.656a1 1 0 0 1-1.415 1.415L7.172 8.586a1 1 0 0 1 0-1.414Z"/></g>`),
		g.Group(children),
	)
}