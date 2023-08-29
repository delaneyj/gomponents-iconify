package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodWinkTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 21a9 9 0 1 1 0-18a9 9 0 0 1 0 18zM9 10h-.01"/><path d="M14.5 15a3.5 3.5 0 0 1-5 0m6-6.5L14 10l1.5 1.5"/></g>`),
		g.Group(children),
	)
}