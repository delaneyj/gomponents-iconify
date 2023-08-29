package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopSunCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M13 17a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0-6a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z" clip-rule="evenodd"/><path d="M3.5 13a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2h-2a1 1 0 0 1-1-1Zm15 0a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1ZM17 17a1 1 0 0 1 1.414 0l1.414 1.414a1 1 0 1 1-1.414 1.414L17 18.414A1 1 0 0 1 17 17ZM6 6a1 1 0 0 1 1.414 0l1.414 1.414a1 1 0 1 1-1.414 1.414L6 7.414A1 1 0 0 1 6 6Zm7 12.5a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0v-2a1 1 0 0 1 1-1Zm0-15a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0v-2a1 1 0 0 1 1-1ZM19.828 6a1 1 0 0 1 0 1.414l-1.414 1.414A1 1 0 1 1 17 7.414L18.414 6a1 1 0 0 1 1.414 0Zm-11 12.414l-1.414 1.414A1 1 0 1 1 6 18.414L7.414 17a1 1 0 0 1 1.414 1.414Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopSunCircleFilled0)"/></g>`),
		g.Group(children),
	)
}