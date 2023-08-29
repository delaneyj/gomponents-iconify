package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidersCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilSlidersCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M6 7.75a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm11.38 0a.5.5 0 0 1 .5-.5h1.62a.5.5 0 0 1 0 1h-1.62a.5.5 0 0 1-.5-.5Z"/><path d="M15.75 9.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM6 17.75a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm11.38 0a.5.5 0 0 1 .5-.5h1.62a.5.5 0 0 1 0 1h-1.62a.5.5 0 0 1-.5-.5Z"/><path d="M15.75 19.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM6 12.75a.5.5 0 0 1 .5-.5h2.13a.5.5 0 0 1 0 1H6.5a.5.5 0 0 1-.5-.5Zm6.5 0a.5.5 0 0 1 .5-.5h6.5a.5.5 0 0 1 0 1H13a.5.5 0 0 1-.5-.5Z"/><path d="M10.75 14.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilSlidersCircleFilled0)"/></g>`),
		g.Group(children),
	)
}