package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileBinary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 7.75v-6h5.5l5 5v7.5M1.75 10.8h3v3.5h-3z"/><path d="M7.25 14.2h3m-3-3.5h1.5v3m-1-11.45v5h5.05"/></g>`),
		g.Group(children),
	)
}