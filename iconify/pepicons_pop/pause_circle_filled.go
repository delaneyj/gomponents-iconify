package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopPauseCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M9.75 6a2 2 0 0 0-2 2v10a2 2 0 1 0 4 0V8a2 2 0 0 0-2-2Zm6.5 0a2 2 0 0 0-2 2v10a2 2 0 1 0 4 0V8a2 2 0 0 0-2-2Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopPauseCircleFilled0)"/></g>`),
		g.Group(children),
	)
}