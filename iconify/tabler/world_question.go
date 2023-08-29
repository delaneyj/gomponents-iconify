package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorldQuestion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.975 11.33a9 9 0 1 0-5.673 9.043M3.6 9h16.8M3.6 15h9.9"/><path d="M11.5 3a17 17 0 0 0 0 18m1-18a16.988 16.988 0 0 1 2.57 9.518m-1.056 5.403A17 17 0 0 1 12.5 21m6.5 1v.01M19 19a2.003 2.003 0 0 0 .914-3.782a1.98 1.98 0 0 0-2.414.483"/></g>`),
		g.Group(children),
	)
}