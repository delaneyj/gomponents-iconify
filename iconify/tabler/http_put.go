package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HttpPut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h2a2 2 0 1 0 0-4H3v8m14-8h4m-2 0v8m-9-8v6a2 2 0 1 0 4 0V8"/>`),
		g.Group(children),
	)
}