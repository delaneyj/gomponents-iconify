package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidersCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopSlidersCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M5.5 7.75a1 1 0 0 1 1-1H13a1 1 0 1 1 0 2H6.5a1 1 0 0 1-1-1Zm11.375 0a1 1 0 0 1 1-1H19.5a1 1 0 1 1 0 2h-1.625a1 1 0 0 1-1-1Z"/><path d="M15.75 8.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm0 2a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5Zm-10.25 7a1 1 0 0 1 1-1H13a1 1 0 1 1 0 2H6.5a1 1 0 0 1-1-1Zm11.375 0a1 1 0 0 1 1-1H19.5a1 1 0 1 1 0 2h-1.625a1 1 0 0 1-1-1Z"/><path d="M15.75 18.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm0 2a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5Zm-10.25-8a1 1 0 0 1 1-1h1.625a1 1 0 0 1 0 2H6.5a1 1 0 0 1-1-1Zm6.5 0a1 1 0 0 1 1-1h6.5a1 1 0 1 1 0 2H13a1 1 0 0 1-1-1Z"/><path d="M10.75 13.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm0 2a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopSlidersCircleFilled0)"/></g>`),
		g.Group(children),
	)
}