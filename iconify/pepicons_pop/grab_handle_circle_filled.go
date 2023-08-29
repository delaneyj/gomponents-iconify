package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrabHandleCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopGrabHandleCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" d="M5.5 11a1 1 0 0 1 0-2h15a1 1 0 1 1 0 2h-15Zm0 3.25a1 1 0 1 1 0-2h15a1 1 0 1 1 0 2h-15Zm0 3.25a1 1 0 1 1 0-2h15a1 1 0 1 1 0 2h-15Z"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopGrabHandleCircleFilled0)"/></g>`),
		g.Group(children),
	)
}