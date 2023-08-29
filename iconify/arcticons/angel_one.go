package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngelOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m8.857 41.446l3.358-5.815H5.5l3.357 5.815z"/><path d="m15.457 41.446l3.357-5.815h-6.715l3.358 5.815zm6.715 0l3.357-5.815h-6.715l3.358 5.815zm6.715 0l3.357-5.815h-6.715l3.358 5.815z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.099 35.631l3.358-5.816H8.742l3.357 5.816zm6.715 0l3.358-5.816h-6.715l3.357 5.816zm6.715 0l3.358-5.816h-6.715l3.357 5.816zm-10.072-5.816L18.814 24h-6.715l3.358 5.815zm6.715 0L25.529 24h-6.715l3.358 5.815zM18.814 24l3.358-5.815h-6.715L18.814 24z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.814 12.369l16.788 29.077H42.5L22.172 6.554l-3.358 5.815z"/>`),
		g.Group(children),
	)
}