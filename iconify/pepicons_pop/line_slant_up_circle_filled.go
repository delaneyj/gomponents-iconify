package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSlantUpCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopLineSlantUpCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M3.808 22.192a1 1 0 0 1 0-1.414l16.97-16.97a1 1 0 1 1 1.414 1.414l-16.97 16.97a1 1 0 0 1-1.414 0Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopLineSlantUpCircleFilled0)"/></g>`),
		g.Group(children),
	)
}