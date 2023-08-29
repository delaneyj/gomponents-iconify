package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13 9a4 4 0 1 0 0 8a4 4 0 0 0 0-8Zm-5 4a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/><path d="M13 11a.5.5 0 0 1 .5.5V13a.5.5 0 0 1-1 0v-1.5a.5.5 0 0 1 .5-.5Z"/><path d="M15 13a.5.5 0 0 1-.5.5H13a.5.5 0 0 1 0-1h1.5a.5.5 0 0 1 .5.5Zm-4-7.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z"/><path d="M11.637 5.02a.5.5 0 0 1 .344.617l-1 3.5a.5.5 0 0 1-.962-.274l1-3.5a.5.5 0 0 1 .618-.344Zm0 15.96a.5.5 0 0 0 .344-.617l-1-3.5a.5.5 0 0 0-.962.274l1 3.5a.5.5 0 0 0 .618.344Zm2.726 0a.5.5 0 0 1-.344-.617l1-3.5a.5.5 0 0 1 .962.274l-1 3.5a.5.5 0 0 1-.618.344Zm0-15.96a.5.5 0 0 0-.344.617l1 3.5a.5.5 0 0 0 .962-.274l-1-3.5a.5.5 0 0 0-.618-.344Z"/><path d="M11 20.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}