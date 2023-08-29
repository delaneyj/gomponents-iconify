package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownLeftCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopDownLeftCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.707 19.707a1 1 0 0 1-1.414 0l-4-4a1 1 0 0 1 0-1.414l4-4a1 1 0 0 1 1.414 1.414L7.414 15l3.293 3.293a1 1 0 0 1 0 1.414Z"/><path d="M15.75 14c.595 0 1.166-.238 1.588-.663a2.284 2.284 0 0 0 .662-1.61V6a1 1 0 1 1 2 0v5.727a4.285 4.285 0 0 1-1.242 3.02A4.239 4.239 0 0 1 15.75 16H6a1 1 0 1 1 0-2h9.75Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopDownLeftCircleFilled0)"/></g>`),
		g.Group(children),
	)
}