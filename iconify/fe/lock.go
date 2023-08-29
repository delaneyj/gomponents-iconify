package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feLock0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLock1" fill="currentColor"><path id="feLock2" d="M7 10V7a5 5 0 1 1 10 0v3h1a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h1Zm-1 2v8h12v-8H6Zm3-2h6V7a3 3 0 0 0-6 0v3Zm5 4h2v4h-2v-4Z"/></g></g>`),
		g.Group(children),
	)
}