package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilWatchCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M13 9a4 4 0 1 0 0 8a4 4 0 0 0 0-8Zm-5 4a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/><path d="M13 11a.5.5 0 0 1 .5.5V13a.5.5 0 0 1-1 0v-1.5a.5.5 0 0 1 .5-.5Z"/><path d="M15 13a.5.5 0 0 1-.5.5H13a.5.5 0 0 1 0-1h1.5a.5.5 0 0 1 .5.5Zm-4-7.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z"/><path d="M11.637 5.02a.5.5 0 0 1 .344.617l-1 3.5a.5.5 0 0 1-.962-.274l1-3.5a.5.5 0 0 1 .618-.344Zm0 15.96a.5.5 0 0 0 .344-.617l-1-3.5a.5.5 0 0 0-.962.274l1 3.5a.5.5 0 0 0 .618.344Zm2.726 0a.5.5 0 0 1-.344-.617l1-3.5a.5.5 0 0 1 .962.274l-1 3.5a.5.5 0 0 1-.618.344Zm0-15.96a.5.5 0 0 0-.344.617l1 3.5a.5.5 0 0 0 .962-.274l-1-3.5a.5.5 0 0 0-.618-.344Z"/><path d="M11 20.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilWatchCircleFilled0)"/></g>`),
		g.Group(children),
	)
}