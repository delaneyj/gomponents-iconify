package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Person(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6 6a4 4 0 1 0 8 0a4 4 0 0 0-8 0Zm6 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z" clip-rule="evenodd"/><path d="M17 18a1 1 0 1 1-2 0v-2.5c0-2.494-2.206-4.5-4.984-4.5C7.23 11 5 13.013 5 15.5l.002 2.5a1 1 0 1 1-2 0L3 15.5C3 11.86 6.169 9 10.016 9C13.86 9 17 11.857 17 15.5V18Z"/></g>`),
		g.Group(children),
	)
}