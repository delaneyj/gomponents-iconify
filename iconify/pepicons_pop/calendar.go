package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3 4h14a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1Zm1 4v8h12V8H4Z" clip-rule="evenodd"/><circle cx="6.5" cy="10.5" r="1.5"/><circle cx="5.5" cy="4.5" r="1.5"/><circle cx="14.5" cy="4.5" r="1.5"/></g>`),
		g.Group(children),
	)
}