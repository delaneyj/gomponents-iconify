package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodSick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 21a9 9 0 1 1 0-18a9 9 0 0 1 0 18zM9 10h-.01M15 10h-.01"/><path d="m8 16l1-1l1.5 1l1.5-1l1.5 1l1.5-1l1 1"/></g>`),
		g.Group(children),
	)
}