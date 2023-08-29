package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodCry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 10h.01M15 10h.01M9.5 15.25a3.5 3.5 0 0 1 5 0m3.066 2.356a2 2 0 1 0 2.897.03L19 16l-1.434 1.606z"/><path d="M20.865 13.517A8.937 8.937 0 0 0 21 12a9 9 0 1 0-9 9c.69 0 1.36-.076 2-.222"/></g>`),
		g.Group(children),
	)
}