package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feUnlock0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUnlock1" fill="currentColor"><path id="feUnlock2" d="M7 10V7a5 5 0 1 1 10 0c0 .55-.45 1-1 1s-1-.45-1-1a3 3 0 0 0-6 0v3h9a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h1Zm-1 2v8h12v-8H6Zm8 2h2v4h-2v-4Z"/></g></g>`),
		g.Group(children),
	)
}