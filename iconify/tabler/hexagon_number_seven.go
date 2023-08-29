package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonNumberSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.02 6.858a2 2 0 0 1 1 1.752v6.555c0 .728-.395 1.4-1.032 1.753l-6.017 3.844a2 2 0 0 1-1.948 0l-6.016-3.844a2 2 0 0 1-1.032-1.752V8.61c0-.728.395-1.4 1.032-1.753l6.017-3.582a2.062 2.062 0 0 1 2 0l6.017 3.583h-.029z"/><path d="M10 8h4l-2 8"/></g>`),
		g.Group(children),
	)
}