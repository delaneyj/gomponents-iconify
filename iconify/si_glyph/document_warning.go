package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.031.031v3.938h3.871L10.031.031z"/><path d="M8.994 0H3.042v16h11V5l-5.048.064V0zM4.958 13l3.48-7l3.52 7h-7z"/><path d="M8 8v2.015h.977V8H8zm0 3h.875v1H8z"/></g>`),
		g.Group(children),
	)
}