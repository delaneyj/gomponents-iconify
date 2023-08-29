package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSpinCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopArrowSpinCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M16.937 7.211a1 1 0 0 1-1.126 1.653A5 5 0 1 0 18 13a1 1 0 1 1 2 0a7 7 0 1 1-3.063-5.789Z"/><path d="M16.538 15.506a1 1 0 1 1-1.077-1.685l3.482-2.227a1 1 0 0 1 1.077 1.685l-3.482 2.227Z"/><path d="M21.903 15.41a1 1 0 0 1-1.826.815l-1.508-3.38a1 1 0 1 1 1.826-.815l1.508 3.38Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopArrowSpinCircleFilled0)"/></g>`),
		g.Group(children),
	)
}