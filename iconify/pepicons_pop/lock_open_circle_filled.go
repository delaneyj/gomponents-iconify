package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpenCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopLockOpenCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M16 10h-6a4 4 0 0 0-4 4v4a4 4 0 0 0 4 4h6a4 4 0 0 0 4-4v-4a4 4 0 0 0-4-4Zm-8 4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-4Z" clip-rule="evenodd"/><path d="M11 11a1 1 0 1 1-2 0V7a4 4 0 1 1 8 0v.5a1 1 0 1 1-2 0V7a2 2 0 1 0-4 0v4Z"/><path fill-rule="evenodd" d="M11 16a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm2.5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopLockOpenCircleFilled0)"/></g>`),
		g.Group(children),
	)
}