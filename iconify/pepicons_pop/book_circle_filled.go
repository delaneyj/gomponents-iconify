package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopBookCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M14 8.297a1 1 0 0 0-.838-.987L5.323 6.026A2 2 0 0 0 3 8v9.737a2 2 0 0 0 1.69 1.976l8.155 1.275A1 1 0 0 0 14 20V8.297Zm-9 9.44V8l7 1.147v9.684l-7-1.094Z"/><path d="M23 8a2 2 0 0 0-2.323-1.974L12.838 7.31a1 1 0 0 0-.838.987V20a1 1 0 0 0 1.155.988l8.154-1.276A2 2 0 0 0 23 17.737V8Zm-2 9.737l-7 1.094V9.147L21 8v9.737Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopBookCircleFilled0)"/></g>`),
		g.Group(children),
	)
}