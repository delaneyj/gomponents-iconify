package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pizza(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 21.5c-3.04 0-5.952-.714-8.5-1.983L12 3l8.5 16.517A19.09 19.09 0 0 1 12 21.5z"/><path d="M5.38 15.866a14.94 14.94 0 0 0 6.815 1.634a14.944 14.944 0 0 0 6.502-1.479M13 11.01V11m-2 3v-.01"/></g>`),
		g.Group(children),
	)
}