package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagStraight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="M4 4.5a1 1 0 0 1 1-1h13.08a1 1 0 0 1 .819 1.573L16.6 8.357a.25.25 0 0 0 0 .286l2.299 3.284a1 1 0 0 1-.82 1.573H5a1 1 0 0 1-1-1v-8Z"/><path fill-rule="evenodd" d="M5 4.5a1 1 0 0 1 1 1v13a1 1 0 1 1-2 0v-13a1 1 0 0 1 1-1Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M16.58 3H3.5v8h13.08l-2.3-3.283a1.25 1.25 0 0 1 0-1.434L16.58 3ZM3.5 2a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h13.08a1 1 0 0 0 .819-1.573L15.1 7.143a.25.25 0 0 1 0-.286l2.3-3.284A1 1 0 0 0 16.579 2H3.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3 3a.5.5 0 0 1 .5.5v14a.5.5 0 0 1-1 0v-14A.5.5 0 0 1 3 3Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}