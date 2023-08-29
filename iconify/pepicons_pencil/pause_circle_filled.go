package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPauseCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M11 17V9a1 1 0 0 0-2 0v8a1 1 0 1 0 2 0ZM10 7a2 2 0 0 0-2 2v8a2 2 0 1 0 4 0V9a2 2 0 0 0-2-2Zm7 10V9a1 1 0 1 0-2 0v8a1 1 0 1 0 2 0ZM16 7a2 2 0 0 0-2 2v8a2 2 0 1 0 4 0V9a2 2 0 0 0-2-2Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPauseCircleFilled0)"/></g>`),
		g.Group(children),
	)
}