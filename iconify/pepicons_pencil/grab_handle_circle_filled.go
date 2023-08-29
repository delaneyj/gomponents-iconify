package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrabHandleCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilGrabHandleCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" d="M5.5 10a.5.5 0 0 1 0-1h15a.5.5 0 0 1 0 1h-15Zm0 3.25a.5.5 0 0 1 0-1h15a.5.5 0 0 1 0 1h-15Zm0 3.25a.5.5 0 0 1 0-1h15a.5.5 0 0 1 0 1h-15Z"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilGrabHandleCircleFilled0)"/></g>`),
		g.Group(children),
	)
}