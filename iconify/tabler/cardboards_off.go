package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardboardsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.96 16.953c.026-.147.04-.298.04-.453V8a2 2 0 0 0-2-2h-9M6 6H5a2 2 0 0 0-2 2v8.5A2.5 2.5 0 0 0 5.5 19h1.06a3 3 0 0 0 2.34-1.13l1.54-1.92a2 2 0 0 1 3.12 0l1.54 1.92A3 3 0 0 0 17.44 19h1.06c.155 0 .307-.014.454-.041"/><path d="M7 12a1 1 0 1 0 2 0a1 1 0 1 0-2 0m9.714.7a1 1 0 0 0-1.417-1.411l1.417 1.41zM3 3l18 18"/></g>`),
		g.Group(children),
	)
}