package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopPersonCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M9 9a4 4 0 1 0 8 0a4 4 0 0 0-8 0Zm6 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z" clip-rule="evenodd"/><path d="M20 21a1 1 0 1 1-2 0v-2.5c0-2.494-2.206-4.5-4.984-4.5C10.23 14 8 16.013 8 18.5l.002 2.5a1 1 0 1 1-2 0L6 18.5c0-3.64 3.169-6.5 7.016-6.5C16.86 12 20 14.857 20 18.5V21Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopPersonCircleFilled0)"/></g>`),
		g.Group(children),
	)
}