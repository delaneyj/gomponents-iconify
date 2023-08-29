package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CandleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M14 10h-4a2 2 0 0 0-2 2v9a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-9a2 2 0 0 0-2-2zm-2.746-7.666l-1.55 1.737C8.662 5.348 8.806 7.168 10 8.237a3 3 0 0 0 4.196-4.28l-1.452-1.624a1 1 0 0 0-1.491.001z"/></g>`),
		g.Group(children),
	)
}