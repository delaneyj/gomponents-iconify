package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopListCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M9.5 9a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/><path fill-rule="evenodd" d="M10.5 9a1 1 0 0 1 1-1h7a1 1 0 1 1 0 2h-7a1 1 0 0 1-1-1Zm0 4a1 1 0 0 1 1-1h7a1 1 0 1 1 0 2h-7a1 1 0 0 1-1-1Zm0 4a1 1 0 0 1 1-1h7a1 1 0 1 1 0 2h-7a1 1 0 0 1-1-1Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopListCircleFilled0)"/></g>`),
		g.Group(children),
	)
}