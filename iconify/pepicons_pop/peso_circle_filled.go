package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PesoCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopPesoCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M14 6.5H9v-2h5a5 5 0 0 1 5 5v1a5 5 0 0 1-5 5H9v-2h5a3 3 0 0 0 3-3v-1a3 3 0 0 0-3-3Z"/><path d="M9 4.5a1 1 0 0 1 1 1V21a1 1 0 1 1-2 0V5.5a1 1 0 0 1 1-1Z"/><path d="M5 8.436a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1Zm0 3a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopPesoCircleFilled0)"/></g>`),
		g.Group(children),
	)
}