package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Division(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="M6.5 12.5a1.5 1.5 0 0 1 0-3h10a1.5 1.5 0 0 1 0 3h-10Z"/><circle cx="11" cy="7" r="2"/><circle cx="11" cy="15" r="2"/></g><path d="M5 10.5a.5.5 0 0 1 0-1h10a.5.5 0 0 1 0 1H5Z"/><path fill-rule="evenodd" d="M8.5 6a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0Zm2 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm-2 8a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0Zm2 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}