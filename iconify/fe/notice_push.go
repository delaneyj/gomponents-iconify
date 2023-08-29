package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoticePush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<defs><path id="feNoticePush0" d="M17 11a4 4 0 1 1 0-8a4 4 0 0 1 0 8ZM5 5h6v2H5v12h12v-6h2v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2Z"/></defs><g id="feNoticePush1" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feNoticePush2"><mask id="feNoticePush3" fill="#fff"><use href="#feNoticePush0"/></mask><use id="feNoticePush4" fill="currentColor" fill-rule="nonzero" href="#feNoticePush0"/></g></g>`),
		g.Group(children),
	)
}