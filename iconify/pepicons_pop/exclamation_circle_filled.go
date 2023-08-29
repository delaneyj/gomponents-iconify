package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopExclamationCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M13 5a2 2 0 0 1 2 2v7a2 2 0 1 1-4 0V7a2 2 0 0 1 2-2Z" clip-rule="evenodd"/><path d="M15 19a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopExclamationCircleFilled0)"/></g>`),
		g.Group(children),
	)
}