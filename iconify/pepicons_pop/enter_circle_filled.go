package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopEnterCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.226 10.232a1 1 0 0 1 1.409.128l2.083 2.5a1 1 0 1 1-1.536 1.28l-2.084-2.5a1 1 0 0 1 .128-1.408Z"/><path d="M11.226 16.768a1 1 0 0 1-.128-1.408l2.084-2.5a1 1 0 1 1 1.536 1.28l-2.083 2.5a1 1 0 0 1-1.409.128Z"/><path d="M13.7 13.5a1 1 0 0 1-1 1H6.2a1 1 0 1 1 0-2h6.5a1 1 0 0 1 1 1Zm5.55 7a1 1 0 0 1-1 1h-9a1 1 0 1 1 0-2h9a1 1 0 0 1 1 1Zm0-14a1 1 0 0 1-1 1h-9a1 1 0 1 1 0-2h9a1 1 0 0 1 1 1Z"/><path d="M9.25 21.5a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1Zm0-10a1 1 0 0 1-1-1v-4a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1Zm9.35 10a1 1 0 0 1-1-1v-14a1 1 0 1 1 2 0v14a1 1 0 0 1-1 1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopEnterCircleFilled0)"/></g>`),
		g.Group(children),
	)
}